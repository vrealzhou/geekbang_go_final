package log

import "fmt"

var defaultLogger Logger = &StdoutLogger{}

type Logger interface {
	Logf(level Level, msg string, a ...interface{})
}

type Level int

const (
	DEBUG Level = iota
	INFO
	ERROR
)

var levelNames = []string{"DEBUG", "INFO", "ERROR"}

// A simple implement
type StdoutLogger struct{}

func (l *StdoutLogger) Logf(level Level, msg string, a ...interface{}) {
	fmt.Printf("[%s] %s", levelNames[level], fmt.Sprintf(msg, a...))
}

func Info(msg string) {
	defaultLogger.Logf(INFO, msg)
}

func Infof(msg string, a ...interface{}) {
	defaultLogger.Logf(INFO, msg, a...)
}

func Error(msg string) {
	defaultLogger.Logf(ERROR, msg)
}

func Errorf(msg string, a ...interface{}) {
	defaultLogger.Logf(ERROR, msg, a...)
}
