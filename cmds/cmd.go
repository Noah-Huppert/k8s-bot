package cmds

// Cmd holds the definition of a command. To be registered with a central
// registry
type Cmd struct {
	// Name is the value the user should provide to invoke the command
	Name string

	// Keywords is a list of keyword definitions, for keywords that can be
	// provided with the command
	Keywords []Keyword

	// Args is a list of definitions for arguments which should be provided
	// with the command
	Args []Arg
}

// NewCmd creates and returns a new Cmd. An error is returned if one occurs,
// nil on success
func NewCmd(name string, keywords []Keyword, args []Arg) (Cmd, error) {
	var cmd Cmd

	// Check name not empty
	if len(name) == 0 {
		return cmd, errors.New("name empty, can not be")
	}

	// Create
	cmd = Cmd{
		Name:     name,
		Keywords: keywords,
		Args:     args,
	}

	return cmd, nil
}
