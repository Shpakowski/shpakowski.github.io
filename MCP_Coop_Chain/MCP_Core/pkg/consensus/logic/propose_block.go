package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/consensus"
	"github.com/mcpcoop/chain/pkg/types"
	"time"
)

// ProposeBlock allows a validator to propose a new block for the current round.
// Business rule: Only registered validators can propose. One proposal per round per validator.
// Logs all proposals for monitoring and debugging.
func ProposeBlock(c *consensus.Consensus, block types.Block, proposer string) error {
	for _, v := range c.Validators {
		if v.Address == proposer {
			// Check if already proposed this round
			for _, p := range c.Proposals {
				if p.Round == c.CurrentRound && p.Proposer == proposer {
					return fmt.Errorf("validator %s already proposed in round %d", proposer, c.CurrentRound)
				}
			}
			proposal := consensus.BlockProposal{
				Block:      block,
				Proposer:   proposer,
				Round:      c.CurrentRound,
				ReceivedAt: time.Now(),
			}
			c.Proposals = append(c.Proposals, proposal)
			fmt.Printf("[CONSENSUS] Block proposed by %s in round %d\n", proposer, c.CurrentRound)
			return nil
		}
	}
	return fmt.Errorf("proposer %s is not a registered validator", proposer)
} 