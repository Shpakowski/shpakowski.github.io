// assets.go - Asset management for electronic cooperatives in MCP Coop Chain
// Handles creation, transfer, and queries for assets. All actions are auditable.
package ecoop

import (
	"errors"
	"time"
	"github.com/mcpcoop/chain/pkg/types"
)

// (Asset struct is now defined in .types/types.go)

// AddAsset adds a new asset to the coop. Returns error if symbol exists or archived.
func AddAsset(coop *types.Coop, symbol, name string, totalSupply uint64) error {
	if coop.IsArchived {
		return errors.New("co-op is archived")
	}
	if _, exists := coop.Assets[symbol]; exists {
		return errors.New("asset already exists")
	}
	a := &types.Asset{
		Symbol:      symbol,
		Name:        name,
		TotalSupply: totalSupply,
		Holders:     make(map[string]uint64),
		CreatedAt:   time.Now(),
	}
	coop.Assets[symbol] = a
	RecordEvent(coop, "asset_create", coop.ID, symbol, "system", map[string]string{"name": name, "total_supply": string(totalSupply)})
	return nil
}

// TransferAsset transfers asset from one member to another. Returns error if not enough balance or archived.
func TransferAsset(coop *types.Coop, symbol, from, to string, amount uint64) error {
	if coop.IsArchived {
		return errors.New("co-op is archived")
	}
	a, exists := coop.Assets[symbol]
	if !exists {
		return errors.New("asset not found")
	}
	if a.Holders[from] < amount {
		return errors.New("insufficient balance")
	}
	a.Holders[from] -= amount
	a.Holders[to] += amount
	RecordEvent(coop, "asset_transfer", coop.ID, symbol, from, map[string]string{"to": to, "amount": string(amount)})
	return nil
}

// ListAssets returns all assets in the coop.
func ListAssets(coop *types.Coop) []*types.Asset {
	assets := []*types.Asset{}
	for _, a := range coop.Assets {
		assets = append(assets, a)
	}
	return assets
}

// BurnAsset burns asset tokens from a holder, reducing total supply.
func BurnAsset(coop *types.Coop, symbol string, holder string, amount uint64) error {
	a, ok := coop.Assets[symbol]
	if !ok {
		return errors.New("asset not found")
	}
	if a.Holders[holder] < amount {
		return errors.New("insufficient balance")
	}
	a.Holders[holder] -= amount
	a.TotalSupply -= amount
	RecordEvent(coop, "asset_burn", coop.ID, symbol, holder, map[string]string{"amount": string(amount)})
	return nil
}

// GetBalance returns the balance of a holder for a given asset.
func GetBalance(coop *types.Coop, symbol, holder string) uint64 {
	a, ok := coop.Assets[symbol]
	if !ok {
		return 0
	}
	return a.Holders[holder]
} 