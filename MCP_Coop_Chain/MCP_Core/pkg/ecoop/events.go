// events.go - Event management for electronic cooperatives in MCP Coop Chain
// Handles event creation, logging, and queries. All actions are auditable.
package ecoop

import (
	"time"
	"github.com/mcpcoop/chain/pkg/types"
)

// (Event struct is now defined in .types/types.go)

// NewEvent creates a new event for a coop.
func NewEvent(eventType, coopID, subject, actor string, details map[string]string) types.Event {
	return types.Event{
		Timestamp: time.Now(),
		Type:      eventType,
		CoopID:    coopID,
		Subject:   subject,
		Actor:     actor,
		Details:   details,
	}
} 