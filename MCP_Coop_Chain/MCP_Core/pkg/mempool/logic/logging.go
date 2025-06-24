package logic

import "fmt"

// LogEvent logs mempool or scMempool events (add, remove, error, info).
// Whitepaper: Section 4.1, 4.2
func LogEvent(event, detail string) {
	fmt.Printf("[MEMPOOL] %s: %s\n", event, detail)
} 