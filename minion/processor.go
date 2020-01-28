package minion

import (
	"context"
	"fmt"

	"github.com/fsufitch/underflow/log"
	"github.com/fsufitch/underflow/rpc-spec"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
)

// Processor encapsulates the logic behind the behavior reacting to asynchronously streamed messages from a master
type Processor interface {
	Background(context.Context, <-chan *rpc.MasterMessage, func(*rpc.MinionMessage) error) error
}

// DefaultProcessor is the default implementation of Processor
type DefaultProcessor struct {
	Log           *log.MultiLogger
	StatusHandler StatusHandler
}

// Background receives messages, handles processing, and sends responses
func (p DefaultProcessor) Background(ctx context.Context, masterMessageChan <-chan *rpc.MasterMessage, sendMessage func(*rpc.MinionMessage) error) error {
	select {
	case message, ok := <-masterMessageChan:
		if !ok {
			return errors.New("master message channel closed unexpectedly")
		}
		p.muxMessage(message, sendMessage)

	}

	_ = rpc.MinionMessage_MinionStatus{}
	return nil
}

func (p DefaultProcessor) muxMessage(message *rpc.MasterMessage, sendMessage func(*rpc.MinionMessage) error) error {
	switch message.MessageType.(type) {
	case *rpc.MasterMessage_CheckStatus_:
		status := p.StatusHandler.GetStatus()
		statusMessage := rpc.MinionMessage_MinionStatus{
			Timestamp: &timestamp.Timestamp{
				Seconds: status.time.Unix(),
				Nanos:   int32(status.time.Nanosecond()),
			},
			TotalCapacity: uint32(status.capacity.Total),
			BusyCapacity:  status.capacity.Busy,
		}
		m := &rpc.MinionMessage{
			MessageType: &rpc.MinionMessage_MinionStatus_{MinionStatus: &statusMessage},
		}
		return sendMessage(m)

	default:
		return fmt.Errorf("unrecognized message type: %v", message.MessageType)
	}
}
