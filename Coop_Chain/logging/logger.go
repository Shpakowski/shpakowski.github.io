package logging

import (
	"io"
	"log"
	"log/slog"
	"os"
)

var (
	Logger *slog.Logger
)

// InitLogger инициализирует логгер с выводом в stdout и файл, JSON-формат, ротация 10 МБ
func InitLogger(logLevel string, logPath string) {
	// Открываем файл для логов (создаём, если нет)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Не удалось открыть лог-файл: %v", err)
	}

	// TODO: добавить ротацию по 10 МБ (упрощённо: вручную проверять размер)
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Уровень логирования
	var level slog.Level
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	Logger = slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: level}))
}

// Примеры использования:
// logging.Logger.Info("Запуск узла", "dataDir", "/data")
// logging.Logger.Error("Ошибка", "err", err)
