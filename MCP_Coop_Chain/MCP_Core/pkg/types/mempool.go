// mempool.go - Mempool types for MCP Coop Chain
package types

// txMempool holds pending standard transactions for block inclusion.
type TxMempool struct {
	Pending []Transaction
}

// scMempool holds pending smart contract transactions for block inclusion.
type ScMempool struct {
	Pending []Transaction
}

// Config holds all tunable mempool settings.
type MempoolConfig struct {
	MaxTxCount int    // Maximum number of transactions in mempool
	MaxTxSize  int    // Maximum size of a single transaction (bytes)
	TTL        int    // Time-to-live for transactions (seconds)
} 