package cli

import (
	"context"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn       *grpc.ClientConn
	connConfig = &ConnectionConfig{
		Host: "localhost",
		Port: 8080,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	}
)

type ConnectionConfig struct {
	Host        string
	Port        uint16
	DialOptions []grpc.DialOption
}

func SetConnection(c *grpc.ClientConn) {
	conn = c
}

type ConnectionOpt func(config *ConnectionConfig)

func WithHost(host string) ConnectionOpt {
	return func(config *ConnectionConfig) {
		config.Host = host
	}
}

func WithPort(port uint16) ConnectionOpt {
	return func(config *ConnectionConfig) {
		config.Port = port
	}
}

func SetDialOptions(opts ...grpc.DialOption) ConnectionOpt {
	return func(config *ConnectionConfig) {
		config.DialOptions = opts
	}
}

func Connection(ctx context.Context, opts ...ConnectionOpt) *grpc.ClientConn {
	if conn != nil {
		return conn
	}

	for _, opt := range opts {
		opt(connConfig)
	}

	var err error
	target := connConfig.Host
	if connConfig.Port > 0 {
		target = net.JoinHostPort(connConfig.Host, strconv.Itoa(int(connConfig.Port)))
	}
	conn, err = grpc.DialContext(ctx, target, connConfig.DialOptions...)
	if err != nil {
		Logger().ErrorContext(ctx, "dial failed", "cause", err)
		os.Exit(1)
	}
	return conn
}
