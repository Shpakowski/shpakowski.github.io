// logger.go - Centralized structured logger using slog
package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	Logger *slog.Logger
	once   sync.Once
)

// Init initializes the global logger. Call once at app startup.
func Init() {
	once.Do(func() {
		Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		slog.SetDefault(Logger)
	})
}

// GetLogger returns the global logger instance.
func GetLogger() *slog.Logger {
	if Logger == nil {
		Init()
	}
	return Logger
} 