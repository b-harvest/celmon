package grpc

import (
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
) 

type Client struct {
	host string
	secureConnection bool
	conn *grpc.ClientConn
	tmServiceClient tmservice.ServiceClient
}
