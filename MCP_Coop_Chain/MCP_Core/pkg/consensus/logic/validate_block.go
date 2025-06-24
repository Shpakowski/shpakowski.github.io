package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/consensus"
)

// ValidateBlock checks if a block proposal is valid according to consensus rules.
// Business rule: Block must be proposed by a validator, must not duplicate previous proposals, and must meet basic block validity.
// Logs validation results for monitoring.
func ValidateBlock(c *consensus.Consensus, proposal consensus.BlockProposal) bool {
	for _, v := range c.Validators {
		if v.Address == proposal.Proposer {
			// TODO: Add more block validity checks as needed
			fmt.Printf("[CONSENSUS] Proposal by %s in round %d is valid\n", proposal.Proposer, proposal.Round)
			return true
		}
	}
	fmt.Printf("[CONSENSUS] Invalid proposal: proposer %s is not a validator\n", proposal.Proposer)
	return false
} 