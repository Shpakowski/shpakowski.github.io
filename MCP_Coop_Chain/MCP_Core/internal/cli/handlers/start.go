package handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mcpcoop/chain/internal/cli"
	"github.com/mcpcoop/chain/internal/logger"
	"github.com/mcpcoop/chain/pkg/chain"
	"github.com/mcpcoop/chain/pkg/types"
)

// HandleStart starts the blockchain node and block timer.
func HandleStart(_ *types.Chain) {
	cli.GlobalNodeState.Mu.Lock()
	defer cli.GlobalNodeState.Mu.Unlock()
	if cli.GlobalNodeState.Running {
		fmt.Println("Node already running.")
		logger.Logger.Warn("Start called but node already running")
		return
	}
	c := chain.NewChain()
	err := chain.Load(c, "State/State.json")
	if err != nil {
		fmt.Printf("[ERROR] Failed to load chain: %v\n", err)
		logger.Logger.Error("Failed to load chain", "error", err)
		return
	}
	cli.GlobalNodeState.Chain = c
	cli.GlobalNodeState.Running = true
	cli.GlobalNodeState.StartTime = time.Now()
	cli.GlobalNodeState.BlockTimer = time.NewTicker(1 * time.Second)
	cli.GlobalNodeState.TimerStop = make(chan struct{})
	logger.Logger.Info("Node started", "start_time", cli.GlobalNodeState.StartTime, "height", chain.Height(c))
	fmt.Printf("[INFO] Node started | time: %s | height: %d\n", cli.GlobalNodeState.StartTime.Format(time.RFC3339), chain.Height(c))
	go func() {
		for {
			select {
			case <-cli.GlobalNodeState.BlockTimer.C:
				// Block creation logic here (use in-memory state only)
			case <-cli.GlobalNodeState.TimerStop:
				cli.GlobalNodeState.BlockTimer.Stop()
				return
			}
		}
	}()
	fmt.Println("[INFO] Node is running. Press Ctrl+C to stop.")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	HandleStop(nil)
}
