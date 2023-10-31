package app

import (
	"celmon/clinet/consensus/grpc"
	"context"
)

func (c *Config) getConsNodeHeight(ctx context.Context) (uint64, error) {
	client := grpc.New(c.General.ConsGrpc, c.General.GrpcSecureConnection)
	err := client.Connect(ctx)
	defer client.Terminate(ctx)
	if err != nil {
		return 0, err
	}

	height, err := client.GetLatestBlockHeight(ctx)
	if err != nil {
		return 0, err
	}

	// For compare
	// convert int64 to uint64
	return uint64(height), nil
}
