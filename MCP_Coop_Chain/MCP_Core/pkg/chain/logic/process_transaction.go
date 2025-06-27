package logic

import (
	"errors"

	"github.com/mcpcoop/chain/pkg/crypto"
	"github.com/mcpcoop/chain/pkg/types"
)

func getPubKeyForAddress(c *types.Chain, addr types.Address) ([]byte, error) {
	for _, w := range c.Wallets {
		if w.Address == addr {
			return w.PubKey, nil
		}
	}
	return nil, errors.New("public key not found for address")
}

// ProcessTransaction validates and adds a transaction to the mempool.
// Returns an error if the transaction is invalid or the sender has insufficient balance.
func ProcessTransaction(c *types.Chain, tx types.Transaction) error {
	if tx.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	if c.Balances[string(tx.From)] < tx.Amount {
		return errors.New("insufficient balance")
	}
	pubKey, err := getPubKeyForAddress(c, tx.From)
	if err != nil {
		return errors.New("invalid transaction sender: public key not found")
	}
	if !crypto.VerifySignature(pubKey, []byte(tx.Hash()), []byte(tx.Signature)) {
		return errors.New("invalid transaction signature")
	}
	// Additional validation can be added here
	c.Mempool = append(c.Mempool, tx)
	return nil
}
