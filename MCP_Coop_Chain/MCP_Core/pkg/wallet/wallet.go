// wallet.go - Wallet struct and basic constructors for MCP Coop Chain
// Provides user wallet management (see whitepaper section 2.1, 2.2)
package wallet

import (
	"crypto/ed25519"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/crypto"
)

// Wallet wraps the shared types.Wallet and adds private key and metadata (not serialized)
type Wallet struct {
	types.Wallet
	PrivKey ed25519.PrivateKey `json:"-"` // Never serialize private key
	Meta    *Metadata           `json:"-"` // Optional metadata (tags, alias, etc)
}

// NewRandomWallet creates a new wallet with a random keypair (see whitepaper 2.1)
func NewRandomWallet() (*Wallet, error) {
	keypair, err := crypto.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	address := crypto.AddressFromPubKey(keypair.PublicKey)
	return &Wallet{
		Wallet: types.Wallet{
			Address: types.Address(address),
			PubKey:  keypair.PublicKey,
			Balance: 0,
		},
		PrivKey: keypair.PrivateKey,
	}, nil
}

// NewWalletFromSeed creates a wallet from a 12-word mnemonic seed (see whitepaper 2.2)
func NewWalletFromSeed(seed string) (*Wallet, error) {
	seedBytes, err := crypto.MnemonicToSeed(seed)
	if err != nil {
		return nil, err
	}
	keypair, err := crypto.GenerateKeyPairFromSeed(seedBytes)
	if err != nil {
		return nil, err
	}
	address := crypto.AddressFromPubKey(keypair.PublicKey)
	return &Wallet{
		Wallet: types.Wallet{
			Address: types.Address(address),
			PubKey:  keypair.PublicKey,
			Balance: 0,
		},
		PrivKey: keypair.PrivateKey,
	}, nil
} 