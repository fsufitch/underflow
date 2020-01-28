package runtime

import (
	"context"

	"github.com/fsufitch/underflow/log"
)

// CLIRunFunc is a plain function that runs an application in a CLI environment
type CLIRunFunc func() error

// BlockingServerRunFunc is a function running a blocking server, meant to be wrapped in a CLI environment by CLIRunFunc
type BlockingServerRunFunc func(context.Context) error

// ProvideCLIRunFunc creates an ApplicationRunFunc that runs a webserver and stops on interrupt
func ProvideCLIRunFunc(logger *log.MultiLogger, runServer BlockingServerRunFunc, interrupt InterruptChannel) CLIRunFunc {
	return func() error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		errChan := make(chan error)

		go func() {
			errChan <- runServer(ctx)
		}()

		select {
		case <-interrupt:
			logger.Infof("interrupt received, shutting down")
			cancel()
		case err := <-errChan:
			logger.Criticalf("fatal server error: %v\n", err)
			return err
		}
		return nil
	}
}
