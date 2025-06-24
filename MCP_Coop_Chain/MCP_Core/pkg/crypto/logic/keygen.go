package logic

import (
	"crypto/ed25519"
	"crypto/rand"
	"bytes"
	"fmt"
)

// GenerateKeyPair generates a new random Ed25519 key pair.
// Business requirement: Secure, random key generation for wallet creation (whitepaper section 2.1).
// Logs key generation for monitoring.
func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("[CRYPTO] Error generating key pair: %v\n", err)
		return nil, nil, err
	}
	fmt.Printf("[CRYPTO] Generated new key pair\n")
	return pub, priv, nil
}

// GenerateKeyPairFromSeed generates a key pair from a seed (e.g., mnemonic-derived).
// Business requirement: Deterministic key generation for wallet recovery (whitepaper section 2.2).
func GenerateKeyPairFromSeed(seed []byte) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(bytes.NewReader(seed))
	if err != nil {
		fmt.Printf("[CRYPTO] Error generating key pair from seed: %v\n", err)
		return nil, nil, err
	}
	fmt.Printf("[CRYPTO] Generated key pair from seed\n")
	return pub, priv, nil
} 