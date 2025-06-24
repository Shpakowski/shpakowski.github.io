// meta.go - Wallet metadata management for MCP Coop Chain
// Business purpose: User-friendly wallet management (see whitepaper section 2.4)
package wallet

// Metadata holds tags, aliases, and other user-defined info for a wallet
// Not serialized on-chain, only for local management
//
type Metadata struct {
	Alias string   `json:"alias,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	Notes string   `json:"notes,omitempty"`
}

// SetAlias sets the wallet's alias
func (w *Wallet) SetAlias(alias string) {
	if w.Meta == nil {
		w.Meta = &Metadata{}
	}
	w.Meta.Alias = alias
}

// AddTag adds a tag to the wallet
func (w *Wallet) AddTag(tag string) {
	if w.Meta == nil {
		w.Meta = &Metadata{}
	}
	w.Meta.Tags = append(w.Meta.Tags, tag)
}

// SetNotes sets notes for the wallet
func (w *Wallet) SetNotes(notes string) {
	if w.Meta == nil {
		w.Meta = &Metadata{}
	}
	w.Meta.Notes = notes
} 