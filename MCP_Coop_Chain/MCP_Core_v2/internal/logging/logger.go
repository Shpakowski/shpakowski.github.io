package logging

import "go.uber.org/zap"

// Logger — интерфейс централизованного логгера для всех сервисов MCP Coop Chain.
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, err error, fields ...zap.Field)
}
