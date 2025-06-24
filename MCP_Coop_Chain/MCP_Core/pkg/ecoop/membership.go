// membership.go - Member management for electronic cooperatives in MCP Coop Chain
// Handles join/leave, role changes, and member queries. All actions are auditable.
package ecoop

import (
	"errors"
	"time"
	"github.com/mcpcoop/chain/pkg/types"
)

// AddMember adds a new member to the coop. Returns error if already a member or archived.
func AddMember(coop *types.Coop, address, role string, meta map[string]string) error {
	if coop.IsArchived {
		return errors.New("co-op is archived")
	}
	if _, exists := coop.Members[address]; exists {
		return errors.New("member already exists")
	}
	m := &types.Member{
		Address:  address,
		Role:     role,
		JoinedAt: time.Now(),
		Meta:     meta,
	}
	coop.Members[address] = m
	RecordEvent(coop, "member_join", coop.ID, address, address, meta)
	return nil
}

// RemoveMember removes a member from the coop. Returns error if not found or archived.
func RemoveMember(coop *types.Coop, address string) error {
	if coop.IsArchived {
		return errors.New("co-op is archived")
	}
	m, exists := coop.Members[address]
	if !exists {
		return errors.New("member not found")
	}
	now := time.Now()
	m.LeftAt = &now
	RecordEvent(coop, "member_leave", coop.ID, address, address, nil)
	return nil
}

// ChangeMemberRole changes a member's role. Returns error if not found or archived.
func ChangeMemberRole(coop *types.Coop, address, newRole string) error {
	if coop.IsArchived {
		return errors.New("co-op is archived")
	}
	m, exists := coop.Members[address]
	if !exists {
		return errors.New("member not found")
	}
	oldRole := m.Role
	m.Role = newRole
	RecordEvent(coop, "member_role_change", coop.ID, address, address, map[string]string{"old_role": oldRole, "new_role": newRole})
	return nil
}

// IsMember returns true if the address is an active member.
func IsMember(coop *types.Coop, address string) bool {
	m, ok := coop.Members[address]
	return ok && m.LeftAt == nil
}

// ListMembers returns all current members (active only).
func ListMembers(coop *types.Coop) []*types.Member {
	members := []*types.Member{}
	for _, m := range coop.Members {
		if m.LeftAt == nil {
			members = append(members, m)
		}
	}
	return members
} 