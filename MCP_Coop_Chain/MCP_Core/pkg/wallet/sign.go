// sign.go - Transaction signing for MCP Coop Chain wallets
// Business purpose: Transaction authentication and integrity (see whitepaper section 3.1)
package wallet

import (
	"github.com/mcpcoop/chain/pkg/crypto"
)

// SignData signs arbitrary data with the wallet's private key
func (w *Wallet) SignData(data []byte) ([]byte, error) {
	return crypto.SignData(w.PrivKey, data)
}

// VerifySignature verifies a signature for data using the wallet's public key
func (w *Wallet) VerifySignature(data, sig []byte) bool {
	return crypto.VerifySignature(w.PubKey, data, sig)
} 