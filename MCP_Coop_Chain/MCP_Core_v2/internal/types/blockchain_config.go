package types

import (
	"encoding/json"
	"os"
	"time"
)

// BlockchainConfig описывает все параметры сети MCP Coop Chain.
type BlockchainConfig struct {
	Logger LoggerConfig `json:"logger"` // Конфиг логгера
	// TokenName — название основной валюты сети (например, MCP Coin)
	TokenName string `json:"tokenName"`
	// TokenSymbol — символ токена (например, MCP)
	TokenSymbol string `json:"tokenSymbol"`
	// StartEmission — стартовая эмиссия, начисляется первому кошельку
	StartEmission float64 `json:"startEmission"`

	// GenesisBlock — параметры генезис-блока
	GenesisBlock GenesisConfig `json:"genesisBlock"`

	// Chain — параметры блокчейна
	Chain ChainConfig `json:"chain"`

	// Fees — настройки комиссий
	Fees FeeConfig `json:"fees"`

	// Consensus — параметры консенсуса
	Consensus ConsensusConfig `json:"consensus"`

	// ValidatorRating — параметры рейтинга валидаторов
	ValidatorRating ValidatorRatingConfig `json:"validatorRating"`

	// Staking — параметры стейкинга
	Staking StakingConfig `json:"staking"`

	// Liquidity — параметры ликвидности и курса
	Liquidity LiquidityConfig `json:"liquidity"`

	// ProtoAPI — параметры встроенных контрактов
	ProtoAPI ProtoAPIConfig `json:"protoAPI"`

	// REST — настройки REST API
	REST RESTConfig `json:"rest"`

	// P2P — настройки P2P взаимодействия
	P2P P2PConfig `json:"p2p"`

	// StoragePath — путь к файлу snapshot (по умолчанию data/full_state_snapshot.json)
	StoragePath string `json:"storagePath"`
}

// GenesisConfig описывает параметры генезис-блока.
type GenesisConfig struct {
	Index     int    `json:"index"`     // Индекс блока (0)
	Timestamp int64  `json:"timestamp"` // UNIX timestamp
	Hash      string `json:"hash"`      // Хэш генезис-блока
}

// ChainConfig описывает параметры блокчейна.
type ChainConfig struct {
	BlockTimeSeconds         int  `json:"blockTimeSeconds"`         // Время между блоками (сек)
	TxPerBlock               int  `json:"txPerBlock"`               // Кол-во транзакций для досрочного блока
	MaxTxMempool             int  `json:"maxTxMempool"`             // Макс. размер mempool
	MaxScMempool             int  `json:"maxScMempool"`             // Макс. размер пула контрактов
	MinContractDeploySeconds int  `json:"minContractDeploySeconds"` // Мин. интервал между деплоями контрактов
	AllowEmptyBlocks         bool `json:"allowEmptyBlocks"`         // Разрешать пустые блоки
}

// FeeConfig описывает комиссии сети.
type FeeConfig struct {
	CreateOrgFee          float64 `json:"createOrgFee"`          // Создание e-Coop
	TxFee                 float64 `json:"txFee"`                 // Обычная транзакция
	PriorityTxFee         float64 `json:"priorityTxFee"`         // Приоритетная транзакция
	DeployProtoAPIFee     float64 `json:"deployProtoAPIFee"`     // Деплой/обновление Proto API
	CallProtoAPIMethodFee float64 `json:"callProtoAPIMethodFee"` // Вызов метода Proto API
}

// ConsensusConfig описывает параметры консенсуса.
type ConsensusConfig struct {
	MaxValidators int `json:"maxValidators"` // Макс. число валидаторов
	VotesRequired int `json:"votesRequired"` // Требуемое число голосов для блока
}

// ValidatorRatingConfig описывает параметры рейтинга валидаторов.
type ValidatorRatingConfig struct {
	InitialRating int `json:"initialRating"` // Начальный рейтинг
	LowMin        int `json:"lowMin"`
	LowMax        int `json:"lowMax"`
	StandardMin   int `json:"standardMin"`
	StandardMax   int `json:"standardMax"`
	EliteMin      int `json:"eliteMin"`
	EliteMax      int `json:"eliteMax"`
}

// StakingConfig описывает параметры стейкинга.
type StakingConfig struct {
	RequiredStake float64 `json:"requiredStake"` // Сумма стейкинга для e-Coop
}

// LiquidityConfig описывает параметры ликвидности и курса.
type LiquidityConfig struct {
	TargetRate                    float64 `json:"targetRate"`                    // Целевой курс MCP/USDC
	BurnFeeThreshold              float64 `json:"burnFeeThreshold"`              // Порог для сжигания комиссий
	ProfitRedistributionThreshold float64 `json:"profitRedistributionThreshold"` // Порог для перераспределения прибыли
}

// ProtoAPIConfig описывает параметры встроенных контрактов.
type ProtoAPIConfig struct {
	ContractIDs      []string `json:"contractIDs"` // Идентификаторы встроенных контрактов
	MaxCallsPerBlock int      `json:"maxCallsPerBlock"`
	MaxGasPerBlock   int      `json:"maxGasPerBlock"`
}

// RESTConfig описывает настройки REST API.
type RESTConfig struct {
	Port    int           `json:"port"`    // Порт сервера
	Timeout time.Duration `json:"timeout"` // Таймаут HTTP-запросов
}

// P2PConfig описывает настройки P2P.
type P2PConfig struct {
	MaxPeers int    `json:"maxPeers"` // Макс. число peer-соединений
	Port     int    `json:"port"`     // Порт
	Protocol string `json:"protocol"`
}

// LoadBlockchainConfigFromFile загружает конфигурацию блокчейна из JSON-файла
// path — путь к файлу конфигурации (например, "configs/blockchain_config.json")
func LoadBlockchainConfigFromFile(path string) (*BlockchainConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cfg BlockchainConfig
	dec := json.NewDecoder(file)
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
