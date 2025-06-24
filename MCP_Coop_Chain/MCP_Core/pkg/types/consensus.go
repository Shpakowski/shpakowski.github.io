// consensus.go - Consensus types for MCP Coop Chain
package types

import (
	"crypto/ed25519"
	"time"
)

// Validator represents a consensus validator node
//
type Validator struct {
	Address   string
	PubKey    ed25519.PublicKey
	Rating    int // Reputation or performance score
	Penalties int // Number of penalties
}

// BlockProposal represents a proposed block in a round
//
type BlockProposal struct {
	Block      Block
	Proposer   string // Validator address
	Round      int
	ReceivedAt time.Time
}

// Vote represents a validator's vote for a block proposal
//
type ConsensusVote struct {
	Voter     string // Validator address
	Proposal  BlockProposal
	Signature []byte
	Round     int
}

// Consensus holds all consensus state and configuration
//
type Consensus struct {
	Validators      []Validator
	CurrentRound    int
	Proposals       []BlockProposal
	Votes           []ConsensusVote
	ValidatorRatings map[string]int
	Config          ConsensusConfig
	Timers          map[int]*time.Timer // round -> timer
}

// ConsensusConfig holds all tunable consensus parameters
//
type ConsensusConfig struct {
	QuorumSize      int
	RoundDuration   int
	PenaltyThreshold int
	MaxValidators   int
} 