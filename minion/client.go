package minion

import (
	"context"
	"io"
	"time"

	"github.com/fsufitch/underflow/config"
	"github.com/fsufitch/underflow/log"
	"github.com/fsufitch/underflow/rpc-spec"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client encapsulates an Underflow Minion running in client mode
type Client struct {
	MasterAddr config.UnderflowMasterAddr
	Log        *log.MultiLogger
	Processor  Processor
}

// Run starts an Underflow Minion in blocking client mode, passing any received messages to Processor
func (c Client) Run(ctx context.Context) error {
	conn, err := grpc.Dial(string(c.MasterAddr))
	defer conn.Close()
	if err != nil {
		return errors.Wrap(err, "could not connect to server")
	}
	client := rpc.NewUnderflowMasterServiceClient(conn)

	// Run handshake
	handshakeCtx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()
	resp, err := client.Handshake(handshakeCtx, &rpc.HandshakeRequest{
		// TODO proper handshake request
	})
	if err != nil {
		return errors.Wrap(err, "handshake failed")
	}

	if !resp.Ok {
		return errors.New("handshake returned not ok, aborting")
	}

	stream, err := client.Stream(ctx)
	if err != nil {
		return errors.Wrap(err, "could not open stream")
	}

	procCtx, cancel2 := context.WithCancel(ctx)
	defer cancel2()

	messageChan := make(chan *rpc.MasterMessage, 5)
	go c.Processor.Background(procCtx, messageChan, stream.Send)

streamLoop:
	for {
		message, err := stream.Recv()
		switch err {
		case io.EOF: // Graceful termination
			c.Log.Errorf("Master message stream gracefully terminated")
			break streamLoop
		case nil:
			messageChan <- message
		default:
			c.Log.Errorf("error receiving message from master: %v", err)
		}
	}

	return nil
}
