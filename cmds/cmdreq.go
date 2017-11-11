package cmds

// CmdReq holds information about command requests received by the bot. Values
// populated by Parse()
type CmdReq struct {
	// Augmentations holds the augmentations present in the command request
	Augmentations []string

	// Cmd holds the name of the command specified in the request
	Cmd string

	// Args holds the arguments provided in the command request. Keys are arg
	// names.
	Args map[string]string
}

// NewCmdReq creates and returns a new command request. An error is returned
// if one occurs, nil on success
func NewCmdReq(augmentations []Augment, cmd string, args map[string]string) (CmdReq, error) {
	var cmdReq CmdReq

	// Check cmd not empty
	if len(cmd) == 0 {
		return cmdReq, errors.New("cmd empty, can not be")
	}

	// Create
	cmdReq = CmdReq{
		Augmentations: augmentations,
		Cmd:           cmd,
		Args:          args,
	}

	return cmdReq, nil
}
