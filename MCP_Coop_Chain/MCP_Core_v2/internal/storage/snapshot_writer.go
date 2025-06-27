package storage

import (
	"encoding/json"
	"mcp-coop-chain/internal/types"
	"os"
)

// FlushSnapshot сериализует всё состояние в data/full_state_snapshot.json (замена файла)
func FlushSnapshot(state *types.FullChainSnapshot) error {
	file, err := os.OpenFile("data/full_state_snapshot.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(state)
}
