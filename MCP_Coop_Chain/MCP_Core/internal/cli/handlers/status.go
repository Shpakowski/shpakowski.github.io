package handlers

import (
	"fmt"
	"time"

	"github.com/mcpcoop/chain/internal/cli"
	"github.com/mcpcoop/chain/internal/logger"
	"github.com/mcpcoop/chain/pkg/chain"
)

// HandleStatus prints the current status of the blockchain node using only in-memory state.
func HandleStatus() {
	cli.GlobalNodeState.Mu.Lock()
	defer cli.GlobalNodeState.Mu.Unlock()
	if !cli.GlobalNodeState.Running || cli.GlobalNodeState.Chain == nil {
		fmt.Println("Node status: stopped")
		fmt.Println("No live metrics available.")
		logger.Logger.Warn("Status requested while node stopped")
		return
	}
	uptime := time.Since(cli.GlobalNodeState.StartTime)
	height := chain.Height(cli.GlobalNodeState.Chain)
	mempool := len(cli.GlobalNodeState.Chain.Mempool)
	fmt.Println("Node status: running")
	fmt.Printf("Block height: %d\n", height)
	fmt.Printf("Mempool transactions: %d\n", mempool)
	fmt.Printf("Uptime: %s\n", formatUptime(uptime))
	logger.Logger.Info("Status requested", "height", height, "mempool", mempool, "uptime", uptime)
}

func formatUptime(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}
