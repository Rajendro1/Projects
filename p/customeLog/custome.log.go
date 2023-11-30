package customlog

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

// LogLevel represents the log level.
type LogLevel int

const (
	// Info level
	Info LogLevel = iota
	// Warning level
	Warning
	// Error level
	Error
)

var (
	logMutex sync.Mutex
)

// logWithDetails is a private function to log messages with additional details.
func logWithDetails(level LogLevel, format string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()

	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()

	// Construct the log message using fmt.Sprintf
	logMessage := fmt.Sprintf("[%s] %s:%d %s --> %s", levelToString(level), file, line, funcName, fmt.Sprintf(format, args...))

	// Use Goroutine for concurrent logging
	go func() {
		log.Print(logMessage)
	}()
}

// levelToString converts LogLevel to a string representation.
func levelToString(level LogLevel) string {
	switch level {
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Info logs an informational message.
func Info1(format string, args ...interface{}) {
	logWithDetails(Info, format, args...)
}

// Warning logs a warning message.
func Warning1(format string, args ...interface{}) {
	logWithDetails(Warning, format, args...)
}

// Error logs an error message.
func Error1(format string, args ...interface{}) {
	logWithDetails(Error, format, args...)
}
