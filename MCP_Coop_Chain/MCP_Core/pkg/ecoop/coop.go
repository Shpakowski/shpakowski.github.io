// coop.go - Core cooperative (Coop) structure and management logic for MCP Coop Chain
// Provides creation, metadata, member/asset access, and serialization helpers.
// All changes are auditable and integrate with blockchain events.
package ecoop

import (
	"time"
	"github.com/google/uuid"
	"github.com/mcpcoop/chain/pkg/types"
)

// (Coop struct is now defined in .types/types.go)

// NewCoop creates a new cooperative with a unique ID and empty members/assets.
func NewCoop(name, description string) *types.Coop {
	return &types.Coop{
		ID:          uuid.NewString(),
		Name:        name,
		CreatedAt:   time.Now(),
		Description: description,
		Members:     make(map[string]*types.Member),
		Assets:      make(map[string]*types.Asset),
		Rules:       DefaultRules(),
		IsArchived:  false,
		Events:      []types.Event{},
	}
}

// UpdateMetadata updates the coop's description.
func UpdateMetadata(c *types.Coop, description string) {
	c.Description = description
	RecordEvent(c, "coop_metadata_update", c.ID, "coop", "system", map[string]string{"description": description})
}

// ArchiveCoop marks the coop as archived (cannot be modified further).
func ArchiveCoop(c *types.Coop) {
	c.IsArchived = true
	RecordEvent(c, "coop_archived", c.ID, "coop", "system", nil)
}

// GetMember returns a member by address, if present.
func GetMember(c *types.Coop, address string) (*types.Member, bool) {
	m, ok := c.Members[address]
	return m, ok
}

// GetAsset returns an asset by symbol, if present.
func GetAsset(c *types.Coop, symbol string) (*types.Asset, bool) {
	a, ok := c.Assets[symbol]
	return a, ok
}

// RecordEvent appends an event to the coop's event log.
func RecordEvent(c *types.Coop, eventType, coopID, subject, actor string, details map[string]string) {
	e := types.Event{
		Timestamp: time.Now(),
		Type:      eventType,
		CoopID:    coopID,
		Subject:   subject,
		Actor:     actor,
		Details:   details,
	}
	c.Events = append(c.Events, e)
}

// ListEvents returns all events for this coop.
func ListEvents(c *types.Coop) []types.Event {
	return c.Events
}

// Serialization helpers (example, can be extended as needed)
// MarshalJSON, UnmarshalJSON, etc. can be implemented if custom logic is needed. 