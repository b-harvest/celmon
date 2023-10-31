package app

import (
	"celmon/clinet/da/rpc"
	"context"
)

func (c *Config) getDANodeHeight(ctx context.Context) (uint64, error) {
	client, err := rpc.New(c.General.DaAPI, c.General.APIToken)
	if err != nil {
		return 0, err
	}

	err = client.Start(ctx)
	if err != nil {
		return 0, err
	}
	defer client.Terminate(ctx)

	height, err := client.GetSyncHeight(ctx)
	if err != nil {
		return 0, err
	}

	return height, nil
}
