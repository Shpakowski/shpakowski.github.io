package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mcpcoop/chain/pkg/chain"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/wallet"
)

// Send creates a new transaction to send coins from one wallet to another
func Send(c *types.Chain, args []string) {
	if len(args) != 3 {
		fmt.Printf("[ERROR] Usage: send <from_address> <to_address> <amount>\n")
		return
	}
	fromAddr := args[0]
	toAddr := args[1]
	amount, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Printf("[ERROR] Invalid amount\n")
		return
	}
	if c.Balances[fromAddr] < amount {
		fmt.Printf("[ERROR] Insufficient balance\n")
		return
	}
	w, err := wallet.GetWalletByAddress(c, fromAddr)
	if err != nil {
		fmt.Printf("[ERROR] Wallet not found or locked\n")
		return
	}
	tx := types.Transaction{
		From:      types.Address(fromAddr),
		To:        types.Address(toAddr),
		Amount:    amount,
		Timestamp: time.Now(),
	}
	if err := w.SignTransaction(&tx); err != nil {
		fmt.Printf("[ERROR] Failed to sign transaction: %v\n", err)
		return
	}
	if err := chain.ProcessTransaction(c, tx); err != nil {
		fmt.Printf("[ERROR] Failed to process transaction: %v\n", err)
		return
	}
	fmt.Printf("[INFO] Transaction added to mempool\nFrom: %s\nTo: %s\nAmount: %.2f\n", fromAddr, toAddr, amount)
}
