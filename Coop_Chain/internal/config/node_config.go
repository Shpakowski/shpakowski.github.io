package config

// NodeConfig содержит локальные настройки узла, которые можно менять и перезагружать без остановки цепи
// (например, путь к данным, порт API, ключ валидатора, лимиты ресурсов)
type NodeConfig struct {
	DataDir       string // Путь к директории данных
	APIPort       int    // Порт для API (REST/gRPC)
	ValidatorKey  string // Путь к приватному ключу валидатора
	MemoryLimitMB int    // Лимит памяти (МБ)
	CPULimit      int    // Лимит CPU (процент)
}

// GetDefaultNodeConfig возвращает параметры по умолчанию для локального узла
func GetDefaultNodeConfig() NodeConfig {
	return NodeConfig{
		DataDir:       "./data",
		APIPort:       8080,
		ValidatorKey:  "$HOME/.mcp/keys/validator.key",
		MemoryLimitMB: 1024,
		CPULimit:      80,
	}
}
