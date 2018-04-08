package metrics

// Counter implements the metric interface for a data field which only
// increases.
type Counter struct {
	// counter is the Prometheus metric definition
	counter prometheus.Counter

	// help contains text describing the metric, to be returned in the Help()
	// method
	help string
}

// NewCounter creates and returns a Counter
func NewCounter(counter prometheus.Counter, help string) *Counter {
	c := &Counter{}

	c.counter = counter
	c.help = help

	return c
}

// definition implement for Metric interface
func (c Counter) definition() prometheus.Collector {
	return c.counter
}

// Help returns a string describing the metric
func (c Counter) Help() string {
	return c.help
}

// Observe implements the Metric interface, by incrementing the counter
// by the provided value. Incrementing by <= 0 will cause the method to
// immediately error.
func (c Counter) Observe(val float64) error {
	// Check not incrementing by <= 0
	if val <= 0 {
		return fmt.Errorf("counter can not be incremented by a negative "+
			"number, or zero, was: %d", val)
	}

	// Set
	c.counter.Add(val)

	return nil
}
