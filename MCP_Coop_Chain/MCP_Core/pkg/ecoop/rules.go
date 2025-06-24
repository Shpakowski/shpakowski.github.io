// rules.go - Governance and operational rules for electronic cooperatives in MCP Coop Chain
// Handles rule definition, validation, and updates. All actions are auditable.
package ecoop

import (
	"errors"
	"github.com/mcpcoop/chain/pkg/types"
)

// (Rules struct is now defined in .types/types.go)

// DefaultRules returns a default set of rules for a new coop.
func DefaultRules() types.Rules {
	return types.Rules{
		AllowJoin:           true,
		AllowLeave:          true,
		ProposalTypes:       map[string]bool{"general": true, "finance": true},
		VotingEligibility:   map[string]bool{"member": true, "admin": true},
		AssetTransferLimits: map[string]uint64{},
		MinBalances:         map[string]uint64{},
		RolePermissions:     map[string][]string{"admin": {"all"}, "member": {"vote", "propose"}},
		Quorum:              1,
	}
}

// UpdateRules updates the coop's rules. Returns error if archived.
func UpdateRules(coop *types.Coop, newRules types.Rules) error {
	if coop.IsArchived {
		return errors.New("co-op is archived")
	}
	coop.Rules = newRules
	RecordEvent(coop, "rules_update", coop.ID, "rules", "system", nil)
	return nil
}

// ValidateProposalType checks if a proposal type is allowed.
func ValidateProposalType(coop *types.Coop, proposalType string) bool {
	return coop.Rules.ProposalTypes[proposalType]
}

// CanVote checks if a role is eligible to vote.
func CanVote(coop *types.Coop, role string) bool {
	return coop.Rules.VotingEligibility[role]
}

// ValidateJoin checks if an address can join the co-op.
func ValidateJoin(coop *types.Coop, address string) error {
	if !coop.Rules.AllowJoin {
		return errors.New("joining not allowed")
	}
	if _, exists := coop.Members[address]; exists {
		return errors.New("already a member")
	}
	return nil
}

// ValidateLeave checks if an address can leave the co-op.
func ValidateLeave(coop *types.Coop, address string) error {
	if !coop.Rules.AllowLeave {
		return errors.New("leaving not allowed")
	}
	if _, exists := coop.Members[address]; !exists {
		return errors.New("not a member")
	}
	return nil
}

// ValidateProposal checks if a proposal is allowed by rules.
func ValidateProposal(coop *types.Coop, proposal *types.Proposal) error {
	if !coop.Rules.ProposalTypes[proposal.Type] {
		return errors.New("proposal type not allowed")
	}
	return nil
}

// ValidateVote checks if a voter is eligible to vote on a proposal.
func ValidateVote(coop *types.Coop, proposal *types.Proposal, voter string) error {
	m, ok := coop.Members[voter]
	if !ok || m.LeftAt != nil {
		return errors.New("not a member")
	}
	if !coop.Rules.VotingEligibility[m.Role] {
		return errors.New("role not eligible to vote")
	}
	return nil
}

// ValidateAssetTransfer checks if an asset transfer is allowed by rules.
func ValidateAssetTransfer(coop *types.Coop, asset *types.Asset, from, to string, amount uint64) error {
	limit, ok := coop.Rules.AssetTransferLimits[asset.Symbol]
	if ok && amount > limit {
		return errors.New("transfer amount exceeds limit")
	}
	min, ok := coop.Rules.MinBalances[asset.Symbol]
	if ok && asset.Holders[from]-amount < min {
		return errors.New("insufficient balance after transfer")
	}
	return nil
} 