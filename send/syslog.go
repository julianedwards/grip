// +build linux freebsd solaris darwin

package send

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"
	"sync"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type syslogger struct {
	logger   *syslog.Writer
	fallback *log.Logger
	*base
	sync.RWMutex
}

// NewSyslogLogger creates a new Sender object taht writes all
// loggable messages to a syslog instance on the specified
// network. Uses the Go standard library syslog implementation that is
// only available on Unix systems. Use this constructor to return a
// connection to a remote Syslog interface, but will fall back first
// to the local syslog interface before writing messages to standard
// output.
func NewSyslogLogger(name, network, raddr string, l LevelInfo) (Sender, error) {
	s := &syslogger{base: newBase(name)}

	s.reset = func() {
		s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
		w, err := syslog.Dial(network, raddr, syslog.Priority(l.Default), s.Name())
		if err != nil {
			s.fallback.Printf("error restarting syslog [%s] for logger: %s", err.Error(), s.Name())
			return
		}
		s.logger = w
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.reset()
	return s, nil
}

// NewLocalSyslogLogger is a constructor for creating the same kind of
// Sender instance as NewSyslogLogger, except connecting directly to
// the local syslog service. If there is no local syslog service, or
// there are issues connecting to it, writes logging messages to
// standard error.
func NewLocalSyslogLogger(name string, l LevelInfo) (Sender, error) {
	return NewSyslogLogger(name, "", "", l)
}

func (s *syslogger) Close() error     { return s.logger.Close() }
func (s *syslogger) Type() SenderType { return Syslog }

func (s *syslogger) Send(m message.Composer) {
	if s.level.ShouldLog(m) {
		msg := m.Resolve()

		if err := s.sendToSysLog(m.Priority(), msg); err != nil {
			s.fallback.Println("syslog error:", err.Error())
			s.fallback.Printf("[p=%d]: %s\n", m.Priority(), msg)
		}
	}
}

func (s *syslogger) sendToSysLog(p level.Priority, message string) error {
	switch p {
	case level.Emergency:
		return s.logger.Emerg(message)
	case level.Alert:
		return s.logger.Alert(message)
	case level.Critical:
		return s.logger.Crit(message)
	case level.Error:
		return s.logger.Err(message)
	case level.Warning:
		return s.logger.Warning(message)
	case level.Notice:
		return s.logger.Notice(message)
	case level.Info:
		return s.logger.Info(message)
	case level.Debug:
		return s.logger.Debug(message)
	}

	return fmt.Errorf("encountered error trying to send: {%s}. Possibly, priority related", message)
}
