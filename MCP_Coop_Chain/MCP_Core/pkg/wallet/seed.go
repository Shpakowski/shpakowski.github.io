// seed.go - Wallet creation and recovery from seed phrase for MCP Coop Chain
// Business purpose: Deterministic wallet recovery (see whitepaper section 2.2)
package wallet

import (
	"github.com/mcpcoop/chain/pkg/crypto"
)

// GenerateSeedFromMnemonic returns a 32-byte seed from a 12-word mnemonic
func GenerateSeedFromMnemonic(mnemonic string) ([]byte, error) {
	return crypto.MnemonicToSeed(mnemonic)
}

// GenerateMnemonicFromSeed returns a 12-word mnemonic from a 32-byte seed
func GenerateMnemonicFromSeed(seed []byte) (string, error) {
	return crypto.SeedToMnemonic(seed)
} 