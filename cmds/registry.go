package cmds

// Registry holds command definitions and matches received messages with these
// definitions.
type Registry struct {
	// Cmds holds all the registered commands
	Cmds map[string]Cmd

	// Augmentations holds all valid argumentation that can be understood by
	// the bot. Keys are augment names. Values are Keywords.
	Augmentations map[string]Keyword
}

// AddCmd registers a command definition with the registry. Returning an error
// if one occurs, nil otherwise.
func (r Registry) AddCmd(cmd Cmd) error {
	// Check name not empty
	if len(cmd.Name) == 0 {
		return errors.New("command name empty")
	}

	// Check if exists
	if _, ok := r.Cmds[cmd.Name]; ok {
		return fmt.Errorf("command with name \"%s\" already registered", cmd.Name)
	}

	// Add
	r.Cmds[cmd.Name] = cmd

	return nil
}

// AddAugment registers an augmentation with the registry. Returning an error
// if one occurs, nil otherwise
func (r Registery) AddAugment(aug Keyword) error {
	// Check augment name not empty
	if len(aug.Name) == 0 {
		return errors.New("augment name empty, cannot be")
	}

	// Check if exists
	if _, ok := r.Augmentations[aug.Name]; ok {
		return errors.New("augment with name \"%s\" already registered", aug.Name)
	}

	// Add
	r.Augmentations[aug.Name] = aug
}

// ParseText interprets received message text and returns a CmdReq struct
// which indicates the command the user sent.
// An error is returned if one occurs, nil on success.
func (r Registry) ParseText(txt string) (CmdReq, error) {
	var CmdReq cmdReq

	// Check not empty
	if len(txt) == 0 {
		return errors.New("error parsing text, was empty")
	}

	// Split by spaces
	words := strings.Fields(txt)

	// Parse any augmentations before command name
    var augmentations []string
    // TODO: Sort out the data structure for augmentation definitions
    // TODO: Finish parse, rn working on parsing augmentations before main cmd
    for _, ok := r.Augmentations[words[0]]; ok {
        augmentations = append(augmentations)
    }
}
