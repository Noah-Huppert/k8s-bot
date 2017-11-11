package cmds

// Arg holds the definition for a command argument
type Arg struct {
	// Name of argument, to be used in documentation, error message, ect
	Name string

	// Optional indicates if the argument can be left out
	Optional bool
}

// NewArg creates and returns a new arg. An error is returned if one occurs,
// nil on success
func NewArg(name string, optional bool) (Arg, error) {
	var arg Arg

	// Check name not empty
	if len(name) == 0 {
		return arg, errors.New("name empty, can not be")
	}

	// Create
	arg = Arg{
		Name:     name,
		Optional: optional,
	}

	return arg, nil
}
