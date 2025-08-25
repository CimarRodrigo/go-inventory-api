package logger

import "log/slog"

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(l *slog.Logger) *SlogLogger {
	return &SlogLogger{
		logger: l,
	}
}

func (l *SlogLogger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}
