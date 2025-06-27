package transaction

import (
	"crypto/ed25519"
	"encoding/json"
	"mcp-chain/crypto"
	"mcp-chain/types"
)

// Используйте types.Transaction для работы с транзакциями

// Tx — структура транзакции
// Payload зарезервировано для вызова Proto-контрактов
// ID = SHA-256(JSON без поля Sig)
type Tx struct {
	ID      types.Hash    `json:"id"`      // Хэш транзакции
	From    types.Address `json:"from"`    // Отправитель
	To      types.Address `json:"to"`      // Получатель
	Amount  uint64        `json:"amount"`  // Сумма в nanoMCP
	Fee     uint64        `json:"fee"`     // Комиссия
	Nonce   uint64        `json:"nonce"`   // Нонc
	Payload []byte        `json:"payload"` // Для контрактов
	Sig     []byte        `json:"sig"`     // Подпись Ed25519
}

// CalcID вычисляет SHA-256(JSON без поля Sig)
func (tx *Tx) CalcID() types.Hash {
	txCopy := *tx
	txCopy.Sig = nil
	data, _ := json.Marshal(txCopy)
	return types.Hash(crypto.HashBytes(data))
}

// Sign подписывает ID приватным ключом Ed25519
func (tx *Tx) Sign(privKey ed25519.PrivateKey) {
	tx.ID = tx.CalcID()
	tx.Sig = ed25519.Sign(privKey, []byte(tx.ID))
}

// Verify проверяет подпись Ed25519
func (tx *Tx) Verify(pubKey ed25519.PublicKey) bool {
	return ed25519.Verify(pubKey, []byte(tx.ID), tx.Sig)
}
