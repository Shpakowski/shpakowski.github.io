package logging

import "go.uber.org/zap"

// ZapLogger — реализация интерфейса Logger на базе Uber Zap.
type ZapLogger struct {
	logger *zap.Logger
}

func (z *ZapLogger) Debug(msg string, fields ...zap.Field) {
	z.logger.Debug(msg, fields...)
}

func (z *ZapLogger) Info(msg string, fields ...zap.Field) {
	z.logger.Info(msg, fields...)
}

func (z *ZapLogger) Warn(msg string, fields ...zap.Field) {
	z.logger.Warn(msg, fields...)
}

func (z *ZapLogger) Error(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))
	z.logger.Error(msg, fields...)
}
