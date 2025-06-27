package handlers

import (
	"fmt"
	"time"

	"github.com/mcpcoop/chain/internal/cli"
	"github.com/mcpcoop/chain/internal/logger"
	"github.com/mcpcoop/chain/pkg/chain"
)

// HandleStop stops the blockchain node.
func HandleStop(_ interface{}) {
	cli.GlobalNodeState.Mu.Lock()
	defer cli.GlobalNodeState.Mu.Unlock()
	if !cli.GlobalNodeState.Running {
		fmt.Println("Node already stopped.")
		logger.Logger.Warn("Stop called but node already stopped")
		return
	}
	if cli.GlobalNodeState.Chain != nil {
		err := chain.Save(cli.GlobalNodeState.Chain, "State/State.json")
		if err != nil {
			fmt.Printf("[ERROR] Failed to save chain: %v\n", err)
			logger.Logger.Error("Failed to save chain", "error", err)
		}
	}
	if cli.GlobalNodeState.BlockTimer != nil && cli.GlobalNodeState.TimerStop != nil {
		close(cli.GlobalNodeState.TimerStop)
		cli.GlobalNodeState.BlockTimer = nil
		cli.GlobalNodeState.TimerStop = nil
	}
	logger.Logger.Info("Node stopped", "stop_time", time.Now())
	fmt.Printf("[INFO] Node stopped | time: %s\n", time.Now().Format(time.RFC3339))
	cli.GlobalNodeState.Running = false
	cli.GlobalNodeState.Chain = nil
	cli.GlobalNodeState.StartTime = time.Time{}
}
