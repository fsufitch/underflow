package minion

import (
	"context"

	"github.com/fsufitch/underflow/log"
	"github.com/fsufitch/underflow/net"
	"github.com/fsufitch/underflow/rpc-spec"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Server encapsulates an Underflow Minion running in server (reverse) mode
// Implements the rpc.UnderflowMinionServiceServer API
type Server struct {
	Log    *log.MultiLogger
	Listen net.UnderflowListenFunc
}

// ReverseHandshake implements an incoming reverse handshake request
func (s Server) ReverseHandshake(ctx context.Context, r *rpc.ReverseHandshakeRequest) (*rpc.ReverseHandshakeResponse, error) {
	// TODO
	return nil, nil
}

// ReverseHandshakeAck implements an incoming reverse handshake request acknowledgement
func (s Server) ReverseHandshakeAck(ctx context.Context, r *rpc.ReverseHandshakeAckRequest) (*rpc.ReverseHandshakeAckResponse, error) {
	// TODO
	return nil, nil
}

// Stream implements an incoming stream request
func (s Server) Stream(stream rpc.UnderflowMinionService_StreamServer) error {
	// TODO
	return nil
}

// Run starts an Underflow Minion server (reverse mode)
func (s Server) Run(ctx context.Context) error {
	lis, err := s.Listen(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to open Underflow listener")
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterUnderflowMinionServiceServer(grpcServer, s)
	s.Log.Infof("registered Underflow minion server (reverse mode); starting...")
	return grpcServer.Serve(lis)
}
