package commands

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/chain"
)

var (
	ticker     *time.Ticker
	tickerStop chan struct{}
)

// startBlockTimer starts the periodic block creation timer
func startBlockTimer(c *types.Chain) {
	if ticker != nil {
		return
	}
	ticker = time.NewTicker(60 * time.Second)
	tickerStop = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("[TIMER] %s | mempool: %d | height: %d\n", time.Now().Format(time.RFC3339), len(c.Mempool), chain.Height(c))
				if len(c.Mempool) > 0 {
					chain.AddBlock(c, c.Mempool)
					fmt.Printf("[TIMER] Block created! New height: %d\n", chain.Height(c))
				} else {
					fmt.Println("[TIMER] No transactions in mempool, block not created.")
				}
			case <-tickerStop:
				ticker.Stop()
				ticker = nil
				return
			}
		}
	}()
}

// stopBlockTimer stops the block creation timer
func stopBlockTimer() {
	if ticker != nil && tickerStop != nil {
		close(tickerStop)
	}
}

// Start starts the blockchain node
type nodeStatus struct {
	Running  bool
	StartTime time.Time
}

var status = nodeStatus{Running: false}

func Start(c *types.Chain, args []string) {
	if status.Running {
		fmt.Printf("[WARN] Node is already running\n")
		return
	}
	status.Running = true
	status.StartTime = time.Now()
	fmt.Printf("[INFO] Node started | height: %d\n", chain.Height(c))

	startBlockTimer(c)
	fmt.Println("[INFO] Node is running. Press Ctrl+C to stop.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("\n[INFO] Shutting down node...")
	status.Running = false
	stopBlockTimer()
	fmt.Println("[INFO] Node stopped")
} 