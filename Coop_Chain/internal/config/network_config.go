package config

// NetworkConfig содержит параметры, которые нельзя менять без перезапуска всей сети
// (например, ID сети, параметры консенсуса, минимальная комиссия и т.д.)
type NetworkConfig struct {
	NetworkID            string  // Уникальный идентификатор сети
	GenesisHeight        uint64  // Высота генезис-блока (обычно 0)
	GenesisAddress       string  // Адрес для начисления генезис-баланса
	GenesisAmount        float64 // Сумма на GENESIS-адресе
	MinFee               float64 // Минимальная комиссия за транзакцию ($MCP)
	ConsensusAlgo        string  // Алгоритм консенсуса (например, "Proof-of-Cooperation")
	BlockReward          float64 // Награда за блок ($MCP)
	MaxTxPerBlock        int     // Максимум транзакций в блоке
	BlockTimeoutMins     int     // Таймаут (минут) до майнинга блока
	MinStake             float64 // Минимальный стейк для валидатора ($MCP)
	MinValidatorStake    float64 // Минимальный стейк для подписи блока
	MaxDrift             int64   // Максимально допустимый дрейф времени (сек)
	AutoMineInterval     int64   // Интервал авто-майна блока (сек)
	ValidatorAddress     string  // Адрес валидатора для генезиса
	GenesisReward        float64 // Награда GENESIS ($MCP)
	MemPoolLimit         int     // Максимальный размер mempool
	MaxGasPerBlock       uint64  // Максимальный gas для блока контрактов
	WalletCreationReward float64 // Награда за создание кошелька
}

// GetDefaultNetworkConfig возвращает параметры по умолчанию для основной сети
func GetDefaultNetworkConfig() NetworkConfig {
	return NetworkConfig{
		NetworkID:            "mcp-coop-chain",
		GenesisHeight:        0,
		GenesisAddress:       "GENESIS",
		GenesisAmount:        1111.0,
		MinFee:               0.001,
		ConsensusAlgo:        "Proof-of-Cooperation",
		BlockReward:          1.0,
		MaxTxPerBlock:        100,
		BlockTimeoutMins:     10,
		MinStake:             1000.0,
		MinValidatorStake:    1000.0,
		MaxDrift:             60,
		AutoMineInterval:     20,
		ValidatorAddress:     "GENESIS_VALIDATOR",
		GenesisReward:        1111.0,
		MemPoolLimit:         1000,
		MaxGasPerBlock:       1000000,
		WalletCreationReward: 100.0,
	}
}
