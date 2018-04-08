package metrics

// Label records metadata about a metric
type Label struct {
	// Key is the name of the metadata
	Key string

	// Value is the value of the metadata
	Value string
}
