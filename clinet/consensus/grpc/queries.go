package grpc

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

func (c *Client) GetLatestBlockHeight(ctx context.Context) (int64, error) {
	resp, err := c.tmServiceClient.GetLatestBlock(
		ctx,
		&tmservice.GetLatestBlockRequest{
		},
	)
	if err != nil {
		return 0, err
	}
	return resp.SdkBlock.Header.Height, nil
}
