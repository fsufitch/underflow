package master

import (
	"context"
	"errors"

	"github.com/fsufitch/underflow/config"
)

// RunFunc is a function that runs an Underflow Master (in  either client or server mode) in a blocking manner
type RunFunc func(context.Context) error

// ProvideRunFunc creates a RunFunc to execute the proper type of master
func ProvideRunFunc(mode config.UnderflowMode, server Server, client Client) (RunFunc, error) {
	switch mode {
	case config.ModeServer:
		return server.Run, nil
	case config.ModeClient:
		return client.Run, nil
	default:
		return nil, errors.New("unknown underflow mode")
	}
}
