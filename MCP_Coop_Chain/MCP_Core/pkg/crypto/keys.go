// keys.go - Crypto public interfaces and type declarations
package crypto

import (
	"crypto/ed25519"
	"github.com/mcpcoop/chain/pkg/crypto/logic"
)

// KeyPair represents a public/private key pair
//
type KeyPair struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

// --- Public Crypto Interfaces ---

// GenerateKeyPair generates a new random key pair
func GenerateKeyPair() (*KeyPair, error) {
	pub, priv, err := logic.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return &KeyPair{PublicKey: pub, PrivateKey: priv}, nil
}

// GenerateKeyPairFromSeed generates a key pair from a seed (e.g., mnemonic)
func GenerateKeyPairFromSeed(seed []byte) (*KeyPair, error) {
	pub, priv, err := logic.GenerateKeyPairFromSeed(seed)
	if err != nil {
		return nil, err
	}
	return &KeyPair{PublicKey: pub, PrivateKey: priv}, nil
}

// SignData signs data with a private key
func SignData(priv ed25519.PrivateKey, data []byte) ([]byte, error) {
	return logic.SignData(priv, data)
}

// VerifySignature verifies a signature for data and public key
func VerifySignature(pub ed25519.PublicKey, data, sig []byte) bool {
	return logic.VerifySignature(pub, data, sig)
}

// MnemonicToSeed converts a 12-word mnemonic to a seed
func MnemonicToSeed(mnemonic string) ([]byte, error) {
	return logic.MnemonicToSeed(mnemonic)
}

// SeedToMnemonic converts a seed to a 12-word mnemonic
func SeedToMnemonic(seed []byte) (string, error) {
	return logic.SeedToMnemonic(seed)
}

// HashData hashes arbitrary data (for tx, blocks, etc)
func HashData(data []byte) []byte {
	return logic.HashData(data)
}

// AddressFromPubKey derives a wallet address from a public key
func AddressFromPubKey(pub ed25519.PublicKey) string {
	return logic.AddressFromPubKey(pub)
} 