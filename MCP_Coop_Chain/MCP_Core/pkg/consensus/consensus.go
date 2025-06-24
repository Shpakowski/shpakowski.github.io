// consensus.go - Consensus struct, configs, and public interfaces
package consensus

import (
	"crypto/ed25519"
	"github.com/mcpcoop/chain/pkg/types"
	"time"
)

// --- Consensus Configs ---

const (
	// Minimum number of validators required to reach consensus
	ConsensusQuorum = 2
	// Duration of each consensus round (in seconds)
	ConsensusRoundDuration = 10
	// Maximum number of validators allowed
	MaxValidators = 21
	// Penalty threshold for slashing or removal
	PenaltyThreshold = 3
)

// --- Consensus State ---

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
	Block      types.Block
	Proposer   string // Validator address
	Round      int
	ReceivedAt time.Time
}

// Vote represents a validator's vote for a block proposal
//
type Vote struct {
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
	Votes           []Vote
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

// --- Public Consensus Methods (interfaces) ---

// ProposeBlock allows a validator to propose a new block
func (c *Consensus) ProposeBlock(block types.Block) error { return nil }

// ValidateBlock checks if a block proposal is valid
func (c *Consensus) ValidateBlock(proposal BlockProposal) bool { return false }

// CollectVotes collects votes for a proposal in the current round
func (c *Consensus) CollectVotes() error { return nil }

// ApplyPenalties applies penalties to misbehaving validators
func (c *Consensus) ApplyPenalties() error { return nil }

// UpdateRatings updates validator ratings based on performance
func (c *Consensus) UpdateRatings() error { return nil }

// CheckQuorum checks if quorum is reached for the current round
func (c *Consensus) CheckQuorum() bool { return false }

// StartEventLoop starts the consensus event/timer loop
func (c *Consensus) StartEventLoop() {} 