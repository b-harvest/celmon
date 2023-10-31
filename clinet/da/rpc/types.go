package rpc

import (
	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

type Client struct {
	RPCClient *client.Client
	host string
	token string
}
