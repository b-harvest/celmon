package rpc

import (
	"context"
)

func (c *Client) GetSyncHeight(ctx context.Context) (uint64, error) {
	resp, err := c.RPCClient.Header.SyncState(ctx)
	if err != nil {
		return 0, err
	}

	return resp.Height, nil
}
