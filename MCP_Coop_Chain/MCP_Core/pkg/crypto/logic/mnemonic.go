package logic

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// MnemonicToSeed converts a 12-word mnemonic to a 32-byte seed.
// Business requirement: Deterministic wallet recovery from mnemonic (whitepaper section 2.2).
// Logs conversion for monitoring.
func MnemonicToSeed(mnemonic string) ([]byte, error) {
	words := strings.Fields(mnemonic)
	if len(words) != 12 {
		fmt.Printf("[CRYPTO] Invalid mnemonic: must be 12 words\n")
		return nil, fmt.Errorf("mnemonic must be 12 words")
	}
	seed := sha256.Sum256([]byte(mnemonic))
	fmt.Printf("[CRYPTO] Mnemonic converted to seed\n")
	return seed[:], nil
}

// SeedToMnemonic converts a 32-byte seed to a 12-word mnemonic (mock implementation).
// Business requirement: Human-readable backup for wallet recovery (whitepaper section 2.2).
// Logs conversion for monitoring.
func SeedToMnemonic(seed []byte) (string, error) {
	if len(seed) != 32 {
		fmt.Printf("[CRYPTO] Invalid seed length for mnemonic\n")
		return "", fmt.Errorf("seed must be 32 bytes")
	}
	// Mock: just hex-encode for now (replace with BIP39 for production)
	mnemonic := fmt.Sprintf("%x", seed)[:48] // 12 words, 4 chars each (mock)
	fmt.Printf("[CRYPTO] Seed converted to mnemonic (mock)\n")
	return mnemonic, nil
} 