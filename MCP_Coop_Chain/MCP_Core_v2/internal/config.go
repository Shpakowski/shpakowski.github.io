package internal

import (
	"mcp-coop-chain/internal/types"
	"sync"
)

var (
	CurrentConfig *types.BlockchainConfig
	configOnce    sync.Once
)

// LoadConfig инициализирует глобальный конфиг (однократно)
func LoadConfig(path string) (*types.BlockchainConfig, error) {
	var err error
	configOnce.Do(func() {
		CurrentConfig, err = types.LoadBlockchainConfigFromFile(path)
	})
	return CurrentConfig, err
}
