package log

import (
	"fmt"
	"os"
	"time"
)

type Level uint8

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

const (
	prefixError = " \033[91mERROR\033[0m "
	prefixWarn  = " \033[33mWARN \033[0m "
	prefixInfo  = " \033[32mINFO \033[0m "
	prefixDebug = " \033[36mDEBUG\033[0m "
	prefixTrace = " \033[36mTRACE\033[0m "
)

const (
	timeStampFmt = "[2006-01-02T15:04:05.000Z07:00]"
)

type Logger interface {
	LogLevel(level Level) Logger

	Trace(any ...interface{}) Logger
	Tracef(format string, any ...interface{}) Logger

	Debug(any ...interface{}) Logger
	Debugf(format string, any ...interface{}) Logger

	Info(any ...interface{}) Logger
	Infof(format string, any ...interface{}) Logger

	Warn(any ...interface{}) Logger
	Warnf(format string, any ...interface{}) Logger

	Error(any ...interface{}) Logger
	Errorf(format string, any ...interface{}) Logger
}

func Log() Logger {
	return defLogger
}

var defLogger = &logger{level: LevelInfo}

type logger struct {
	level Level
}

func (l *logger) LogLevel(level Level) Logger {
	l.level = level
	return l
}

func (l *logger) Trace(any ...interface{}) Logger {
	return l.writeln(LevelTrace, any)
}

func (l *logger) Tracef(format string, any ...interface{}) Logger {
	return l.writef(LevelTrace, format, any)
}

func (l *logger) Debug(any ...interface{}) Logger {
	return l.writeln(LevelDebug, any)
}

func (l *logger) Debugf(format string, any ...interface{}) Logger {
	return l.writef(LevelDebug, format, any)
}

func (l *logger) Info(any ...interface{}) Logger {
	return l.writeln(LevelInfo, any)
}

func (l *logger) Infof(format string, any ...interface{}) Logger {
	return l.writef(LevelInfo, format, any)
}

func (l *logger) Warn(any ...interface{}) Logger {
	return l.writeln(LevelWarn, any)
}

func (l *logger) Warnf(format string, any ...interface{}) Logger {
	return l.writef(LevelWarn, format, any)
}

func (l *logger) Error(any ...interface{}) Logger {
	fmt.Print(l.leader(LevelError))
	_, _ = fmt.Fprintln(os.Stderr, any...)
	return l
}

func (l *logger) Errorf(format string, any ...interface{}) Logger {
	fmt.Print(l.leader(LevelError))
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", any...)
	return l
}

func (l *logger) writef(lvl Level, form string, any []interface{}) Logger {
	if lvl >= l.level {
		fmt.Print(l.leader(lvl))
		fmt.Printf(form+"\n", any...)
	}
	return l
}

func (l *logger) writeln(lvl Level, any []interface{}) Logger {
	if lvl >= l.level {
		fmt.Print(l.leader(lvl))
		fmt.Println(any...)
	}
	return l
}

func (l *logger) leader(lvl Level) string {
	return time.Now().Format(timeStampFmt) + lvl.String()
}

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return prefixTrace
	case LevelDebug:
		return prefixDebug
	case LevelInfo:
		return prefixInfo
	case LevelWarn:
		return prefixWarn
	case LevelError:
		return prefixError
	default:
		panic("invalid state")
	}
}
