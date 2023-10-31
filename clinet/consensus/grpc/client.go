package grpc

import (
	"celmon/log"
	"context"
	"crypto/tls"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func New(host string, secureConnection bool) *Client {
	return &Client{
		host: host,
		secureConnection: secureConnection,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	options := []grpc.DialOption{grpc.WithBlock()}
	if c.secureConnection {
		options = append(
			options,
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		)
	} else {
		options = append(
			options,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	}

	conn, err := grpc.DialContext(
		ctx,
		c.host,
		options...,
	)
	if err != nil {
		return err
	}

	c.conn = conn
	c.tmServiceClient = tmservice.NewServiceClient(conn)

	log.Info("GRPC connected")

	return nil
}

func (c *Client) Terminate(_ context.Context) error {
	err := c.conn.Close()
	log.Info("GRPC connection terminated")

	return err
}
