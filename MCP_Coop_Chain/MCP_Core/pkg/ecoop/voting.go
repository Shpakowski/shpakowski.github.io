// voting.go - Proposal and voting logic for electronic cooperatives in MCP Coop Chain
// Handles proposal creation, voting, closing, and result calculation. All actions are auditable.
package ecoop

import (
	"errors"
	"time"
	"github.com/mcpcoop/chain/pkg/types"
)

// (Proposal and Vote structs are now defined in .types/types.go)

// CreateProposal creates a new proposal and adds it to the coop's event log.
func CreateProposal(coop *types.Coop, creator, title, desc, typ string, options []string, deadline time.Time) (*types.Proposal, error) {
	if coop.IsArchived {
		return nil, errors.New("co-op is archived")
	}
	// Only one open proposal per topic
	for _, e := range coop.Events {
		if e.Type == "proposal_create" && e.Details["executed"] != "true" && e.Details["type"] == typ {
			return nil, errors.New("an open proposal of this type already exists")
		}
	}
	id := generateProposalID()
	p := &types.Proposal{
		ID:          id,
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
		Deadline:    deadline,
		Creator:     creator,
		Type:        typ,
		Options:     options,
		Votes:       make(map[string]*types.Vote),
		Executed:    false,
	}
	RecordEvent(coop, "proposal_create", coop.ID, id, creator, map[string]string{"title": title, "type": typ, "executed": "false"})
	return p, nil
}

// CastVote records a member's vote on a proposal.
func CastVote(coop *types.Coop, proposalID, voter, option string) error {
	p, ok := getProposalByID(coop, proposalID)
	if !ok {
		return errors.New("proposal not found")
	}
	if p.Executed {
		return errors.New("proposal already closed")
	}
	if time.Now().After(p.Deadline) {
		return errors.New("voting deadline passed")
	}
	if _, exists := p.Votes[voter]; exists {
		return errors.New("already voted")
	}
	// Option must be valid
	valid := false
	for _, o := range p.Options {
		if o == option {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("invalid voting option")
	}
	p.Votes[voter] = &types.Vote{Voter: voter, Option: option, VotedAt: time.Now()}
	RecordEvent(coop, "vote_cast", coop.ID, proposalID, voter, map[string]string{"option": option})
	return nil
}

// CloseProposal marks a proposal as executed/closed.
func CloseProposal(coop *types.Coop, proposalID string) error {
	p, ok := getProposalByID(coop, proposalID)
	if !ok {
		return errors.New("proposal not found")
	}
	if p.Executed {
		return errors.New("already closed")
	}
	p.Executed = true
	RecordEvent(coop, "proposal_closed", coop.ID, proposalID, "system", nil)
	return nil
}

// GetProposalResult returns the winning option for a proposal.
func GetProposalResult(coop *types.Coop, proposalID string) (string, error) {
	p, ok := getProposalByID(coop, proposalID)
	if !ok {
		return "", errors.New("proposal not found")
	}
	if !p.Executed {
		return "", errors.New("proposal not closed")
	}
	// Tally votes
	counts := make(map[string]int)
	for _, v := range p.Votes {
		counts[v.Option]++
	}
	max := 0
	winner := ""
	for opt, cnt := range counts {
		if cnt > max {
			max = cnt
			winner = opt
		}
	}
	if winner == "" {
		return "", errors.New("no votes cast")
	}
	return winner, nil
}

// Helper: generate a unique proposal ID (could use UUID or increment)
func generateProposalID() string {
	return time.Now().Format("20060102150405")
}

// Helper: get proposal by ID from coop events (for demo, could be improved)
func getProposalByID(c *types.Coop, id string) (*types.Proposal, bool) {
	for _, e := range c.Events {
		if e.Type == "proposal_create" && e.Subject == id {
			return &types.Proposal{ID: id, Executed: e.Details["executed"] == "true"}, true
		}
	}
	return nil, false
} 