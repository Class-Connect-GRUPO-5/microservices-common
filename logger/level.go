package logger

import (
	"fmt"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
	Panic
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	case Panic:
		return "panic"
	default:
		return "unreachable"
	}
}

func LogLevelFromString(s string) (LogLevel, error) {
	switch s {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	case "panic":
		return Panic, nil
	default:
		return -1, fmt.Errorf("invalid level")
	}
}
