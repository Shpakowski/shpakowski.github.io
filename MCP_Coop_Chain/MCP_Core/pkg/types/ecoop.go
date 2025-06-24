// ecoop.go - Electronic cooperative (ecoop) core types for MCP Coop Chain
package types

import "time"

// Coop represents an electronic cooperative on the blockchain.
type Coop struct {
	ID          string                 // Unique coop ID (UUID)
	Name        string                 // Human-readable name
	CreatedAt   time.Time              // Creation timestamp
	Description string                 // Description/metadata
	Members     map[string]*Member     // Members by address
	Assets      map[string]*Asset      // Assets by symbol
	Rules       Rules                  // Governance and operational rules
	IsArchived  bool                   // Archived flag
	Events      []Event                // Event log (in-memory, persisted at chain level)
}

// Member represents a co-op member with role and metadata.
type Member struct {
	Address  string            // Blockchain address
	Role     string            // e.g. "founder", "member", "admin"
	JoinedAt time.Time         // Join timestamp
	LeftAt   *time.Time        // Leave timestamp (nil if active)
	Meta     map[string]string // Display name, tags, etc.
}

// Asset represents a fungible asset/token managed by a co-op.
type Asset struct {
	Symbol      string            // Asset symbol (unique per co-op)
	Name        string            // Human-readable name
	TotalSupply uint64            // Total supply
	Holders     map[string]uint64 // Address to balance
	CreatedAt   time.Time         // Creation timestamp
}

// Proposal represents a governance proposal in the co-op.
type Proposal struct {
	ID          string             // Unique proposal ID
	Title       string             // Title of the proposal
	Description string             // Description/details
	CreatedAt   time.Time          // Creation timestamp
	Deadline    time.Time          // Voting deadline
	Creator     string             // Address of proposal creator
	Type        string             // Proposal type/topic
	Options     []string           // Voting options
	Votes       map[string]*Vote   // Votes by address
	Executed    bool               // Whether proposal has been executed/closed
}

// Vote represents a member's vote on a proposal.
type Vote struct {
	Voter   string    // Address of the voter
	Option  string    // Chosen option
	VotedAt time.Time // Timestamp of vote
}

// Event represents a major action or state change in a co-op.
type Event struct {
	Timestamp time.Time         // When the event occurred
	Type      string            // Event type (e.g. coop_create, member_join, asset_transfer)
	CoopID    string            // Coop ID
	Subject   string            // Main subject (e.g. member address, asset symbol, proposal ID)
	Actor     string            // Who performed the action
	Details   map[string]string // Additional details (optional)
}

// Rules defines governance and operational rules for a co-op.
type Rules struct {
	AllowJoin           bool                // Can new members join?
	AllowLeave          bool                // Can members leave?
	ProposalTypes       map[string]bool     // Allowed proposal types
	VotingEligibility   map[string]bool     // Role eligibility for voting
	AssetTransferLimits map[string]uint64   // Max transfer per tx by asset symbol
	MinBalances         map[string]uint64   // Min balance per asset
	RolePermissions     map[string][]string // Role to allowed actions
	Quorum              int                 // Minimum votes for proposal to pass
} 