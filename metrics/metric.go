package metrics

// Metric is used to collect analytics data from various places in the
// application code.
type Metric interface {
	// definition returns the Prometheus definition of the metric.
	definition() prometheus.Collector

	// Help returns text explaining the metric
	Help() string

	// Observe records the provided metric value. Nil will be return on
	// success, an error on failure.
	Observe(val float64) error
}
