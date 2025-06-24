package logic

import (
	"encoding/json"
	"io/ioutil"
	"github.com/mcpcoop/chain/pkg/types"
)

// LoadChain loads the blockchain state from a JSON file.
func LoadChain(c *types.Chain, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
} 