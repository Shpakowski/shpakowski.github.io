package logic

import (
	"crypto/ed25519"
	"fmt"
)

// SignData signs data with a private key.
// Business requirement: All transactions and blocks must be signed for authenticity (whitepaper section 3.1).
// Logs signing actions for monitoring.
func SignData(priv ed25519.PrivateKey, data []byte) ([]byte, error) {
	sig := ed25519.Sign(priv, data)
	fmt.Printf("[CRYPTO] Data signed\n")
	return sig, nil
}

// VerifySignature verifies a signature for data and public key.
// Business requirement: All signatures must be verifiable by network participants (whitepaper section 3.2).
// Logs verification results for monitoring.
func VerifySignature(pub ed25519.PublicKey, data, sig []byte) bool {
	ok := ed25519.Verify(pub, data, sig)
	if ok {
		fmt.Printf("[CRYPTO] Signature verified\n")
	} else {
		fmt.Printf("[CRYPTO] Signature verification failed\n")
	}
	return ok
} 