package minion

import (
	"context"

	"github.com/fsufitch/underflow/log"
	"github.com/fsufitch/underflow/net"
	"github.com/fsufitch/underflow/rpc-spec"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// RunFunc is a function that runs an Underflow Master server in a blocking manner
type RunFunc func(context.Context) error

// ProvideRunFunc builds a RunFunc to launch an Underflow Master server
func ProvideRunFunc(log *log.MultiLogger, listen net.UnderflowListenFunc, srv rpc.UnderflowMinionServiceServer) RunFunc {
	return func(ctx context.Context) error {
		lis, err := listen(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to open Underflow listener")
		}
		grpcServer := grpc.NewServer()
		rpc.RegisterUnderflowMinionServiceServer(grpcServer, srv)
		log.Infof("registered Underflow minion server; starting...")
		return grpcServer.Serve(lis)
	}
}
