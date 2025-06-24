package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/consensus"
)

// UpdateRatings updates validator ratings based on their performance in the last round.
// Business rule: Validators who propose valid blocks or vote correctly get higher ratings; misbehavior reduces ratings.
// Logs all rating updates for monitoring.
func UpdateRatings(c *consensus.Consensus) error {
	for _, v := range c.Validators {
		// Example: increase rating for participation, decrease for penalties
		if v.Penalties > 0 {
			c.ValidatorRatings[v.Address] -= v.Penalties
			fmt.Printf("[CONSENSUS] Validator %s rating decreased by %d\n", v.Address, v.Penalties)
		} else {
			c.ValidatorRatings[v.Address]++
			fmt.Printf("[CONSENSUS] Validator %s rating increased\n", v.Address)
		}
	}
	return nil
} 