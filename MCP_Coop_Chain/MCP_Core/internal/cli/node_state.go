package cli

import (
	"sync"
	"time"

	"github.com/mcpcoop/chain/pkg/types"
)

type NodeState struct {
	Running    bool
	Chain      *types.Chain
	StartTime  time.Time
	BlockTimer *time.Ticker
	TimerStop  chan struct{}
	Mu         sync.Mutex
}

var GlobalNodeState = &NodeState{}
