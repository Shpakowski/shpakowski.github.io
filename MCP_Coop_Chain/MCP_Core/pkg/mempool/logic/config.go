package logic

// Config holds all tunable mempool settings.
// Whitepaper: Section 4.1, 4.2
// Document each parameter below.
type Config struct {
	MaxTxCount int    // Maximum number of transactions in mempool
	MaxTxSize  int    // Maximum size of a single transaction (bytes)
	TTL        int    // Time-to-live for transactions (seconds)
}

// DefaultConfig returns the default mempool config.
func DefaultConfig() Config {
	return Config{
		MaxTxCount: 1000,
		MaxTxSize:  10240, // 10 KB
		TTL:        3600,  // 1 hour
	}
} 