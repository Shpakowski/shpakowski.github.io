package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/consensus"
)

// CollectVotes collects votes for proposals in the current round.
// Business rule: Each validator can vote once per round. Votes are tallied for each proposal.
// Logs vote collection for monitoring.
func CollectVotes(c *consensus.Consensus) error {
	voteCount := make(map[string]int) // proposal proposer -> count
	for _, vote := range c.Votes {
		if vote.Round == c.CurrentRound {
			voteCount[vote.Proposal.Proposer]++
		}
	}
	for proposer, count := range voteCount {
		fmt.Printf("[CONSENSUS] Proposal by %s received %d votes in round %d\n", proposer, count, c.CurrentRound)
	}
	return nil
} 