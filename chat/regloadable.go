package chat

// RegLoadable wraps the Load method. Which is used to to add command and
// augment definitions.
type RegLoadable interface {
	// Load should be used to add a file's command and or augment definitions
	// to the registry. Nil will be returned on success. An error otherwise.
	Load(registry Registry) error
}
