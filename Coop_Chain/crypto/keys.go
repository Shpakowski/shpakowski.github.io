package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"mcp-chain/types"
	"os"
)

// GenerateKeyPair генерирует новую пару ключей Ed25519
func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	return pub, priv, err
}

// PrivateKeyToHex сериализует приватный ключ в hex-строку
func PrivateKeyToHex(priv ed25519.PrivateKey) string {
	return hex.EncodeToString(priv)
}

// PublicKeyToHex сериализует публичный ключ в hex-строку
func PublicKeyToHex(pub ed25519.PublicKey) string {
	return hex.EncodeToString(pub)
}

// PublicKeyToAddress возвращает адрес (hex) из публичного ключа
func PublicKeyToAddress(pub ed25519.PublicKey) types.Address {
	return types.Address(hex.EncodeToString(pub))
}

// LoadPrivateKeyFromFile загружает приватный ключ из файла
func LoadPrivateKeyFromFile(path string) (ed25519.PrivateKey, error) {
	data, err := ioutil.ReadFile(os.ExpandEnv(path))
	if err != nil {
		return nil, err
	}
	key, err := hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	return ed25519.PrivateKey(key), nil
}

// SavePrivateKeyToFile сохраняет приватный ключ в файл (hex-строка)
func SavePrivateKeyToFile(priv ed25519.PrivateKey, path string) error {
	return ioutil.WriteFile(os.ExpandEnv(path), []byte(PrivateKeyToHex(priv)), 0600)
}
