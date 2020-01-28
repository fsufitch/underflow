package master

import (
	"context"

	"github.com/fsufitch/underflow/log"
	"github.com/fsufitch/underflow/net"
	"github.com/fsufitch/underflow/rpc-spec"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Server encapsulates an Underflow Master running in server mode
// Implements the rpc.UnderflowMasterServiceServer API
type Server struct {
	Log    *log.MultiLogger
	Listen net.UnderflowListenFunc
}

// Handshake implements an incoming handshake request
func (s Server) Handshake(ctx context.Context, r *rpc.HandshakeRequest) (*rpc.HandshakeResponse, error) {
	return nil, nil
}

// Stream implements an incoming stream request
func (s Server) Stream(stream rpc.UnderflowMasterService_StreamServer) error {
	return nil
}

// Run starts an Underflow Master server
func (s Server) Run(ctx context.Context) error {
	lis, err := s.Listen(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to open Underflow listener")
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterUnderflowMasterServiceServer(grpcServer, s)
	s.Log.Infof("registered Underflow master server; starting...")
	return grpcServer.Serve(lis)
}
