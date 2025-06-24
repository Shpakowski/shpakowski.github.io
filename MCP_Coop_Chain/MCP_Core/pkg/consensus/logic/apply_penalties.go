package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/consensus"
)

// ApplyPenalties applies penalties to validators who misbehave (e.g., double voting, not voting, etc).
// Business rule: If a validator exceeds the penalty threshold, they are slashed or removed.
// Logs all penalty actions for monitoring.
func ApplyPenalties(c *consensus.Consensus) error {
	for i, v := range c.Validators {
		if v.Penalties >= c.Config.PenaltyThreshold {
			fmt.Printf("[CONSENSUS] Validator %s exceeded penalty threshold and is removed\n", v.Address)
			// Remove validator (simple removal for demo)
			c.Validators = append(c.Validators[:i], c.Validators[i+1:]...)
		}
	}
	return nil
} 