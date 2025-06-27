package logging

import (
	"mcp-coop-chain/internal/types"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger создает новый логгер по заданной конфигурации.
func NewLogger(cfg types.LoggerConfig) (Logger, error) {
	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}

	zapCfg := zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		Encoding:      "json",
		OutputPaths:   cfg.OutputPaths,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	z, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{logger: z}, nil
}
