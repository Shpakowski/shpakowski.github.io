package types

// LoggerConfig — параметры конфигурации логгера MCP Coop Chain.
type LoggerConfig struct {
	Level       string   // "debug", "info", "warn", "error"
	OutputPaths []string // например: ["stdout", "logs/mcp.log"]
}
