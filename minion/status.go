package minion

import "time"

// Status is a struct containing the status of this client
type Status struct {
	time     time.Time
	capacity Capacity
}

// StatusHandler is an interface that provides the client's current status
type StatusHandler interface {
	GetStatus() Status
}

// DefaultStatusHandler implements StatusHandler
type DefaultStatusHandler struct {
	Capacity *Capacity
}

// GetStatus returns the current status of the client
func (h DefaultStatusHandler) GetStatus() Status {
	capacityCopy := *h.Capacity
	return Status{
		time:     time.Now(),
		capacity: capacityCopy,
	}
}
