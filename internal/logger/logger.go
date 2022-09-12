package logger

type Logger struct {
	name string
}

func NewLogger(name string) *Logger {
	logger := &Logger{name: name}
	return logger
}
