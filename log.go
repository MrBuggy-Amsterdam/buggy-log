package log

import (
	"fmt"
)

// Logging levels
type Level int

const (
	DebugLevel   Level = 0
	InfoLevel    Level = 1
	SuccessLevel Level = 2
	WarningLevel Level = 3
	ErrorLevel   Level = 4
)

var verbosity Level = InfoLevel

func SetVerbosity(level Level) {
	verbosity = level
}

func Debug(format string, a ...any) {
	if verbosity <= DebugLevel {
		var msg = "🟣 DEBUG: " + format + "\n"
		fmt.Printf(msg, a...)
	}
}

func Info(format string, a ...any) {
	if verbosity <= InfoLevel {
		var msg = "🔵 INFO: " + format + "\n"
		fmt.Printf(msg, a...)
	}
}

func Success(format string, a ...any) {
	if verbosity <= SuccessLevel {
		var msg = "🟢 SUCCESS: " + format + "\n"
		fmt.Printf(msg, a...)
	}
}

func Warn(format string, a ...any) {
	if verbosity <= WarningLevel {
		var msg = "🟡 WARN: " + format + "\n"
		fmt.Printf(msg, a...)
	}
}

func Error(format string, a ...any) {
	if verbosity <= ErrorLevel {
		var msg = "🔴 ERROR: " + format + "\n"
		fmt.Printf(msg, a...)
	}
}

func Fatal(format string, a ...any) {
	var msg = "🔴 FATAL: " + format + "\n"
	fmt.Printf(msg, a...)
	panic("A fatal error occurred.")
}
