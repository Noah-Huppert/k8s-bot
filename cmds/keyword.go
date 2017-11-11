package cmds

import "errors"

// Keyword holds information for flag like values which can be included with a
// command.
//
// A keyword has multiple values, one of which can be present at a time. Each
// of these values sets a different value for the same setting.
//
// For example this keyword in the "logs" command. It has the values: "top" and
// "bottom". One of which can be specified at a time. To read the logs either
// from the top or bottom.
type Keyword struct {
	// Name is used to refer to the Keyword in documentation, errors, ect...
	Name string

	// Values holds all the different strings the user can provide to specify
	// this keyword. Each string indicates a different value of the same general
	// setting
	Values []string

	// Optional indicates if the keyword can be left out of a command
	Optional bool
}

// NewKeyword creates and returns a new Keyword from the values provided.
// An error is returned if one occurs, nil otherwise.
func NewKeyword(name string, values []string, optional bool) (*Keyword, error) {
	var keyword Keyword

	// Check if name is empty
	if len(name) == 0 {
		return &keyword, errors.New("Name was empty, can not be")
	}

	// Create
	keyword = Keyword{
		Name:     name,
		Values:   values,
		Optional: optional,
	}

	return &keyword, nil
}
