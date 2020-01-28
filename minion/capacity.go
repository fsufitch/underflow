package minion

// Capacity is a struct containing an abstract measure of a client's capacity
type Capacity struct {
	Total uint64
	Busy  uint64
}

const maxCapacity = 8

// ProvideInitialCapacity creates the initial capacity values for a client
func ProvideInitialCapacity() *Capacity {
	// TODO: smarter configurable capacity
	return &Capacity{
		Total: maxCapacity,
		Busy:  0,
	}
}
