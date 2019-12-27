package logger

import "fmt"

type Logger interface {
	LogError(message string, args ...interface{})
	LogWarning(message string, args ...interface{})
	LogInfo(message string, args ...interface{})
	LogDebug(message string, args ...interface{})
}

type logger struct {
	debug bool
}

func NewLogger(debug bool) *logger {
	return &logger{debug: debug}
}

func (l *logger) LogError(message string, args ...interface{}) {
	fmt.Println(message, args)
}

func (l *logger) LogWarning(message string, args ...interface{}) {
	fmt.Println(message, args)
}

func (l *logger) LogInfo(message string, args ...interface{}) {
	fmt.Println(message, args)
}

func (l *logger) LogDebug(message string, args ...interface{}) {
	if l.debug {
		fmt.Println(message, args)
	}
}
