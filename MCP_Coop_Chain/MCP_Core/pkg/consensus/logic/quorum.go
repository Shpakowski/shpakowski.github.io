package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/consensus"
)

// CheckQuorum checks if the number of votes for any proposal meets the quorum size.
// Business rule: A proposal is accepted if it receives at least quorum votes in the current round.
// Logs quorum checks for monitoring.
func CheckQuorum(c *consensus.Consensus) bool {
	voteCount := make(map[string]int)
	for _, vote := range c.Votes {
		if vote.Round == c.CurrentRound {
			voteCount[vote.Proposal.Proposer]++
		}
	}
	for proposer, count := range voteCount {
		if count >= c.Config.QuorumSize {
			fmt.Printf("[CONSENSUS] Quorum reached for proposal by %s in round %d\n", proposer, c.CurrentRound)
			return true
		}
	}
	fmt.Printf("[CONSENSUS] No proposal reached quorum in round %d\n", c.CurrentRound)
	return false
} 