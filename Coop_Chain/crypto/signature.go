package crypto

import (
	"crypto/ed25519"
)

// Sign подписывает сообщение приватным ключом Ed25519
// Возвращает сырой 64-байтовый срез подписи
func Sign(priv ed25519.PrivateKey, msg []byte) []byte {
	return ed25519.Sign(priv, msg)
}

// Verify проверяет подпись Ed25519
// Возвращает true, если подпись корректна
func Verify(pub ed25519.PublicKey, msg, sig []byte) bool {
	return ed25519.Verify(pub, msg, sig)
}
