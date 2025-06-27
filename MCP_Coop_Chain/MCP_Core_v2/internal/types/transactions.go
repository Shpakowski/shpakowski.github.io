package types

import "time"

// Transaction — универсальная структура транзакции MCP Coop Chain.
// Включает защиту от повторной отправки (nonce), комиссию, подпись, уникальный ID, все бизнес-правила.
type Transaction struct {
	TxID                string    `json:"txId"`                // Уникальный идентификатор транзакции (SHA256)
	Origin              string    `json:"origin"`              // Источник: "User" или "Oracle"
	Payload             []byte    `json:"payload"`             // Основное содержимое (сериализованное)
	SmartContractHashes []string  `json:"smartContractHashes"` // Список затронутых смарт-контрактов
	Fee                 float64   `json:"fee"`                 // Комиссия в MCP
	Reason              string    `json:"reason,omitempty"`    // Причина/комментарий (опционально)
	Timestamp           time.Time `json:"timestamp"`           // Время создания (UTC)
	Signature           string    `json:"signature"`           // Подпись отправителя (ECDSA, base64)
	Nonce               uint64    `json:"nonce"`               // Уникальный номер для защиты от повторной отправки (replay protection)
}
