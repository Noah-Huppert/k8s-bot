package cmds

// RegLoadable wraps the Load method. Which is used to by the RegistryLoader
// to add command and augment definitions.
type RegLoadable interface {
	// Load is called by the RegistryLoader. And should be used to add a file's
	// command and or augment definitions to the registry. Nil will be returned
	// on success. An error otherwise
	Load(registry Registry) error
}
