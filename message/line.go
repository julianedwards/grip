package message

import (
	"fmt"
	"strings"

	"github.com/tychoish/grip/level"
)

type lineMessenger struct {
	Lines   []interface{} `yaml:"lines" json:"lines" bson:"lines"`
	Base    `bson:"metadata" json:"metadata" yaml:"metadata"`
	message string
}

// NewLineMessage is a basic constructor for a type that, given a
// bunch of arguments, calls fmt.Sprintln() on the arguments passed to
// the constructor during the String() operation. Use in combination
// with Compose[*] logging methods.
func NewLineMessage(p level.Priority, args ...interface{}) Composer {
	m := &lineMessenger{
		Lines: args,
	}
	_ = m.SetPriority(p)
	return m
}

// NewLine returns a message Composer roughly equivalent to
// fmt.Sprintln().
func NewLine(args ...interface{}) Composer {
	return &lineMessenger{
		Lines: args,
	}
}

func (l *lineMessenger) Loggable() bool {
	return len(l.Lines) > 0
}

func (l *lineMessenger) String() string {
	if l.message == "" {
		l.message = strings.Trim(fmt.Sprintln(l.Lines...), "\n ")
	}

	return l.message
}

func (l *lineMessenger) Raw() interface{} {
	_ = l.Collect()
	return l
}
