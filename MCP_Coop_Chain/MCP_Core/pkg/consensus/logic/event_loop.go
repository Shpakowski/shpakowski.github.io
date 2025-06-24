package logic

import (
	"fmt"
	"time"
	"github.com/mcpcoop/chain/pkg/consensus"
)

// StartEventLoop starts the consensus event/timer loop.
// Business rule: Each round is timed; at the end of the round, votes are tallied, penalties applied, ratings updated, and next round started.
// Logs all major events for monitoring and debugging.
func StartEventLoop(c *consensus.Consensus) {
	for {
		fmt.Printf("[CONSENSUS] Starting round %d\n", c.CurrentRound)
		t := time.NewTimer(time.Duration(c.Config.RoundDuration) * time.Second)
		c.Timers[c.CurrentRound] = t
		<-t.C
		fmt.Printf("[CONSENSUS] Round %d ended\n", c.CurrentRound)
		// End-of-round processing (collect votes, apply penalties, update ratings, check quorum)
		_ = CollectVotes(c)
		_ = ApplyPenalties(c)
		_ = UpdateRatings(c)
		_ = CheckQuorum(c)
		c.CurrentRound++
	}
} 