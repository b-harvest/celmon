package rpc

import (
	"celmon/log"
	"context"

	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

func New(host string, token string) (*Client, error) {
	result := &Client{
		host: host,
		token: token,
	}

	return result, nil
}

func (c *Client) Start(ctx context.Context) error {
	var err error
	c.RPCClient, err = client.NewClient(ctx, c.host, c.token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Terminate(_ context.Context) {
	c.RPCClient.Close()

	log.Info("RPC connection terminated")
}
