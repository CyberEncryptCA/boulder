// Package log implements a wrapper around the Go standard library's
// logging package. Clients should set the current log level; only
// messages below that level will actually be logged. For example, if
// Level is set to LevelWarning, only log messages at the Warning,
// Error, and Critical levels will be logged.
package log

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
)

// The following constants represent logging levels in increasing levels of seriousness.
const (
	// LevelDebug is the log level for Debug statements.
	LevelDebug = iota
	// LevelInfo is the log level for Info statements.
	LevelInfo
	// LevelWarning is the log level for Warning statements.
	LevelWarning
	// LevelError is the log level for Error statements.
	LevelError
	// LevelCritical is the log level for Critical statements.
	LevelCritical
	// LevelFatal is the log level for Fatal statements.
	LevelFatal
)

var levelPrefix = [...]string{
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
	LevelWarning:  "WARNING",
	LevelError:    "ERROR",
	LevelCritical: "CRITICAL",
	LevelFatal:    "FATAL",
}

var (
	// Level stores the current logging level.
	Level = LevelInfo
	// SysLogger is a syslog Writer to be used if not nil.
	SysLogger *syslog.Writer
)

func init() {
	flag.IntVar(&Level, "loglevel", LevelInfo, "Log level (0 = DEBUG, 5 = FATAL)")
}

func print(l int, msg string) {
	if l >= Level {
		if SysLogger != nil {
			var err error
			switch l {
			case LevelDebug:
				err = SysLogger.Debug(msg)
			case LevelInfo:
				err = SysLogger.Info(msg)
			case LevelWarning:
				err = SysLogger.Warning(msg)
			case LevelError:
				err = SysLogger.Err(msg)
			case LevelCritical:
				err = SysLogger.Crit(msg)
			case LevelFatal:
				err = SysLogger.Emerg(msg)
			}
			if err != nil {
				log.Printf("Unable to write syslog: %v for msg: %s\n", err, msg)
			}
		} else {
			log.Printf("[%s] %s", levelPrefix[l], msg)
		}
	}
}

func outputf(l int, format string, v []interface{}) {
	print(l, fmt.Sprintf(format, v...))
}

func output(l int, v []interface{}) {
	print(l, fmt.Sprint(v...))
}

// Fatalf logs a formatted message at the "fatal" level and then exits. The
// arguments are handled in the same manner as fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	outputf(LevelFatal, format, v)
	os.Exit(1)
}

// Fatal logs its arguments at the "fatal" level and then exits.
func Fatal(v ...interface{}) {
	output(LevelFatal, v)
	os.Exit(1)
}

// Criticalf logs a formatted message at the "critical" level. The
// arguments are handled in the same manner as fmt.Printf.
func Criticalf(format string, v ...interface{}) {
	outputf(LevelCritical, format, v)
}

// Critical logs its arguments at the "critical" level.
func Critical(v ...interface{}) {
	output(LevelCritical, v)
}

// Errorf logs a formatted message at the "error" level. The arguments
// are handled in the same manner as fmt.Printf.
func Errorf(format string, v ...interface{}) {
	outputf(LevelError, format, v)
}

// Error logs its arguments at the "error" level.
func Error(v ...interface{}) {
	output(LevelError, v)
}

// Warningf logs a formatted message at the "warning" level. The
// arguments are handled in the same manner as fmt.Printf.
func Warningf(format string, v ...interface{}) {
	outputf(LevelWarning, format, v)
}

// Warning logs its arguments at the "warning" level.
func Warning(v ...interface{}) {
	output(LevelWarning, v)
}

// Infof logs a formatted message at the "info" level. The arguments
// are handled in the same manner as fmt.Printf.
func Infof(format string, v ...interface{}) {
	outputf(LevelInfo, format, v)
}

// Info logs its arguments at the "info" level.
func Info(v ...interface{}) {
	output(LevelInfo, v)
}

// Debugf logs a formatted message at the "debug" level. The arguments
// are handled in the same manner as fmt.Printf.
func Debugf(format string, v ...interface{}) {
	outputf(LevelDebug, format, v)
}

// Debug logs its arguments at the "debug" level.
func Debug(v ...interface{}) {
	output(LevelDebug, v)
}
