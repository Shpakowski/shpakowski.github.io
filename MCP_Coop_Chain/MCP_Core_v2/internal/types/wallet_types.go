package types

import "time"

// Wallet — публичная часть кошелька MCP Coop Chain (для блокчейна и снапшотов).
type Wallet struct {
	Address   string `json:"address"`   // Адрес кошелька (base58, всегда = SHA256(publicKey))
	PublicKey string `json:"publicKey"` // Публичный ключ (base58, ECDSA P-256)
}

// PrivateWallet — приватная часть кошелька (только для локального хранения).
type PrivateWallet struct {
	PrivateKey string    `json:"privateKey"` // Зашифрованный приватный ключ (base64, AES-GCM)
	PublicKey  string    `json:"publicKey"`  // Публичный ключ (base58, ECDSA P-256)
	Address    string    `json:"address"`    // Адрес кошелька (base58)
	CreatedAt  time.Time `json:"createdAt"`  // Время создания
	Salt       string    `json:"salt"`       // Соль для PBKDF2 (base64)
	AESNonce   string    `json:"aesNonce"`   // Nonce для AES-GCM (base64)
}
