package send

type SenderType int

const (
	Custom SenderType = iota
	Systemd
	Native
	JsonConsole
	JsonFile
	Syslog
	Internal
	File
	Slack
	Xmpp
	Bootstrap
)

func (t SenderType) String() string {
	switch t {
	case Systemd:
		return "systemd"
	case Native:
		return "native"
	case Syslog:
		return "syslog"
	case Internal:
		return "internal"
	case File:
		return "file"
	case Bootstrap:
		return "bootstrap"
	case Custom:
		return "custom"
	case Slack:
		return "slack"
	case Xmpp:
		return "xmpp"
	case JsonConsole:
		return "json-console"
	case JsonFile:
		return "json-console"
	default:
		return "native"
	}
}
