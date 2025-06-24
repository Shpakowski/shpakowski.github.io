package logic

import (
	"encoding/json"
	"io/ioutil"
	"github.com/mcpcoop/chain/pkg/types"
)

// SaveChain saves the blockchain state to a JSON file.
func SaveChain(c *types.Chain, filename string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
} 