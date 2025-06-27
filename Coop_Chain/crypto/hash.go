package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// HashBytes возвращает SHA-256 хэш от произвольного среза байт (в hex-строке)
func HashBytes(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// HashJSON сериализует объект в JSON и возвращает SHA-256 хэш (в hex-строке)
func HashJSON(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return HashBytes(b), nil
}
