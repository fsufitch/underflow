package net

import (
	"context"
	"fmt"
	"net"

	"github.com/fsufitch/underflow/config"
	"github.com/fsufitch/underflow/log"
)

// UnderflowListenFunc is a curried net.Listen wrapper that closes the port on context cancellation
type UnderflowListenFunc func(context.Context) (net.Listener, error)

// ProvideUnderflowListenFunc creates an UnderflowListenFunc for starting an Underflow server
func ProvideUnderflowListenFunc(log *log.MultiLogger, port config.UnderflowListenPort) UnderflowListenFunc {
	return func(ctx context.Context) (net.Listener, error) {
		addr := fmt.Sprintf(":%d", port)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		log.Infof("Underflow listening on %s", addr)
		go func() {
			<-ctx.Done()
			lis.Close()
		}()
		return lis, err
	}
}
