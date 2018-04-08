package metrics

// Metrics is the public interface for setting analytics data. This interface
// will remain as stable as possible, independent of the implementation.
//
// At the time of design Prometheus was used in the implementation. So it is
// highly opinionated. However this extra stable buffer layer will prevent
// the need to rewrite all metric calls throughout the entire codebase.
type Metrics interface {
}

// register initializes the provided metric with the analytics system. Nil is
// returned on success, an error otherwise.
func (m Metrics) register(m metric) error {
	// Register with Prometheus.
	prometheus.MustRegister(m.definition())
}

// Serve starts a http server with an endpoint for Prometheus to scrape. Which
// will be stopped when `ctx` is canceled.
// The
// error that stops execution will be returned. Nil on success.
func (m Metrics) Serve() error {
	return nil
}
