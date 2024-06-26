package driver

import (
	celestia "github.com/ethereum-optimism/optimism/op-celestia"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
)

func SetDAClient(cfg celestia.Config) error {
	client, err := celestia.NewDAClient(cfg.Rpc, cfg.AuthToken, cfg.Namespace, false)
	if err != nil {
		return err
	}
	return derive.SetDAClient(client)
}
