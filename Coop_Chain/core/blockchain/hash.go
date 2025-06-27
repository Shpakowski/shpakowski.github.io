package blockchain

import (
	"mcp-chain/crypto"
	"mcp-chain/types"
)

// CalcBlockHash вычисляет хэш блока (по Header и Body)
func CalcBlockHash(b *types.Block) (types.Hash, error) {
	h, err := crypto.HashJSON(struct {
		Header types.Header
		Body   types.Body
	}{b.Header, b.Body})
	if err != nil {
		return "", err
	}
	return types.Hash(h), nil
}
