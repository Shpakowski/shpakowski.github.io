package handlers

import (
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/wallet"
)

// HandleNewWallet creates a new wallet from a seed phrase and returns the wallet or an error.
func HandleNewWallet(seedPhrase string) (*types.Wallet, error) {
	return wallet.NewWalletFromSeed(seedPhrase)
}
