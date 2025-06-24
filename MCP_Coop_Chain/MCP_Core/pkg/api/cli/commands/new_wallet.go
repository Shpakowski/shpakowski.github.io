package commands

import (
	"fmt"
	"strings"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/wallet"
)

// NewWallet creates a new wallet from a seed phrase using the wallet package
// Business: Deterministic wallet creation and user onboarding (whitepaper 2.2)
func NewWallet(c *types.Chain, args []string) {
	if len(args) == 0 || strings.TrimSpace(args[0]) == "" {
		fmt.Printf("[ERROR] Seed phrase required. Use 'new-wallet <12-word-seed>'\n")
		return
	}
	seedPhrase := strings.Join(strings.Fields(args[0]), " ")
	w, err := wallet.NewWalletFromSeed(seedPhrase)
	if err != nil {
		wallet.LogWalletError("Wallet creation failed", err)
		fmt.Printf("[ERROR] %s\n", err.Error())
		return
	}
	if c.Balances == nil {
		c.Balances = map[string]float64{}
	}
	if _, exists := c.Balances[string(w.Address)]; exists {
		fmt.Printf("Wallet already exists!\nAddress: %s\n", w.Address)
		return
	}
	c.Balances[string(w.Address)] = 100.0 // Give new wallet a mock balance
	wallet.LogWalletEvent("Wallet created", "address", w.Address)
	fmt.Printf("[INFO] New wallet created!\nAddress: %s\nInitial balance: 100.0\n", w.Address)
} 