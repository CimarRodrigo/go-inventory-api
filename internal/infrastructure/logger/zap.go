package logger

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(development bool) (*ZapLogger, error) {
	var l *zap.Logger
	var err error

	if development {

		config := zap.NewDevelopmentConfig()
		config.DisableCaller = false
		config.DisableStacktrace = true

		l, err = config.Build(zap.AddCaller(), zap.AddCallerSkip(2))
	} else {

		config := zap.NewProductionConfig()
		config.DisableCaller = true
		config.DisableStacktrace = true

		l, err = config.Build()
	}

	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: l,
	}, nil
}
func argsToFields(args ...any) []zap.Field {
	fields := []zap.Field{}
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		value := args[i+1]
		fields = append(fields, zap.Any(key, value))
	}
	return fields
}

func (l *ZapLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, argsToFields(args...)...)
}

func (l *ZapLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, argsToFields(args...)...)
}

func (l *ZapLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, argsToFields(args...)...)
}

func (l *ZapLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, argsToFields(args...)...)
}
