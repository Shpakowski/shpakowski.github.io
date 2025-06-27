package types

import (
	"time"
)

// FullChainSnapshot описывает полный снимок состояния сети для отладки и восстановления.
type FullChainSnapshot struct {
	Blocks        []Block        `json:"blocks"`
	Mempool       []Transaction  `json:"mempool"`
	Organizations []Organization `json:"organizations"`
	Wallets       []Wallet       `json:"wallets"`
	State         ChainState     `json:"state"`
	Contracts     []ContractCall `json:"contracts"`
	Timestamp     time.Time      `json:"timestamp"`
}

// Organization — заглушка для MVP.
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ChainState — состояние сети (заглушка для MVP).
type ChainState struct {
	LastHash      string `json:"lastHash"`
	PostStateHash string `json:"postStateHash"`
}
