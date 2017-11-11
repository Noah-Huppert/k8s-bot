package cmds

// RegistryLoader is used to easily add file's command and augment definitions,
// leveraging the RegLoadable interface
type RegistryLoader struct {
	registry Registry
}

// NewRegistryLoader creates and returns a new RegistryLoader instance
func NewRegistryLoader(registry Registry) *RegistryLoader {
	return &RegistryLoader{registry}
}

// Load calls the provided RegLoadable's Load() method
func (r RegistryLoader) Load(loadable RegLoadable) error {
	return loadable.Load(r.registry)
}
