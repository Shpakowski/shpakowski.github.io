package logic

import (
	"time"
	"github.com/mcpcoop/chain/pkg/types"
)

// CreateGenesisBlock returns the genesis block for the blockchain.
func CreateGenesisBlock() types.Block {
	return types.Block{
		BlockHeader: types.BlockHeader{
			Index:     0,
			Timestamp: time.Now(),
			PrevHash:  "",
			Hash:      "genesis",
		},
		Transactions: nil,
	}
}

// NewChain creates a new Chain with a genesis block and empty state.
func NewChain() *types.Chain {
	genesis := CreateGenesisBlock()
	return &types.Chain{
		Blocks:   []types.Block{genesis},
		Balances: map[string]float64{},
		Mempool:  []types.Transaction{},
	}
} 