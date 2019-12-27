package logger

import "fmt"

type Logger interface {
	LogError(message string, resource interface{})
	LogWarning(message string, resource interface{})
	LogInfo(message string, resource interface{})
	LogDebug(message string, resource interface{})
}

type logger struct {
	setting string
}

func NewLogger(setting string) logger {
	return logger{setting: setting}
}

func (l *logger) LogError(message string, resource interface{}) {
	fmt.Println(message, resource)
}

func (l *logger) LogWarning(message string, resource interface{}) {
	fmt.Println(message, resource)
}

func (l *logger) LogInfo(message string, resource interface{}) {
	fmt.Println(message, resource)
}

func (l *logger) LogDebug(message string, resource interface{}) {
	fmt.Println(message, resource)
}
