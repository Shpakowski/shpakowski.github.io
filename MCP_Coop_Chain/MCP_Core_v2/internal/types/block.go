package types

import "time"

// BlockHeader описывает метаданные блока, по которым строится хеш.
type BlockHeader struct {
	BlockID       string    `json:"blockId"`       // Уникальный идентификатор блока (SHA256 от содержимого заголовка)
	Height        uint64    `json:"height"`        // Номер блока в цепочке (начиная с 0)
	Timestamp     time.Time `json:"timestamp"`     // Время создания блока (UTC, формат RFC3339)
	PreviousHash  string    `json:"previousHash"`  // Хеш предыдущего блока
	PostStateHash string    `json:"postStateHash"` // Хеш состояния системы после выполнения этого блока
	Tag           string    `json:"tag,omitempty"` // Тип блока: "" (обычный), "Oracle", "DeployContract"
	TxRoot        string    `json:"txRoot"`        // Меркл-корень всех транзакций в этом блоке
}

// ValidatorSig представляет подпись одного валидатора, подписавшего блок.
type ValidatorSig struct {
	PubKey    string `json:"pubKey"`    // Публичный ключ валидатора (в base58)
	Signature string `json:"signature"` // Подпись заголовка блока (ECDSA, base64)
}

// Block — это основной тип блока в блокчейне MCP Coop Chain.
type Block struct {
	Header        BlockHeader    `json:"header"`        // Заголовок блока
	Proposer      string         `json:"proposer"`      // Публичный ключ/адрес инициатора (создателя) блока
	ValidatorSigs []ValidatorSig `json:"validatorSigs"` // Список подписей валидаторов (до 10)
	Transactions  []Transaction  `json:"transactions"`  // Список универсальных транзакций MCP
	ContractCalls []ContractCall `json:"contractCalls"` // Список вызовов встроенных Proto API контрактов
}
