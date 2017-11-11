package cmds

// CmdReq holds information about command requests received by the bot. Values
// populated by Parse()
type CmdReq struct {
	// RawText holds the unprocessed message text
	RawText string

	// Augments holds the augments present in the command request
	Augments map[string]string

	// Cmd holds the name of the command specified in the request
	Cmd string

	// Args holds the arguments provided in the command request. Keys are arg
	// names.
	Args map[string]string
}

// EmptyCmdReq initialized a CmdReq with blank values
func EmptyCmdReq() CmdReq {
	var cmdReq CmdReq

	cmdReq.Augments = make(map[string]string)
	cmdReq.Args = make(map[string]string)

	return cmdReq
}

// NewCmdReq creates and returns a new command request. An error is returned
// if one occurs, nil on success
func NewCmdReq(augments map[string]string, cmd string, args map[string]string) CmdReq {
	cmdReq := EmptyCmdReq()

	// Create
	cmdReq = CmdReq{
		Augments: augments,
		Cmd:      cmd,
		Args:     args,
	}

	return cmdReq
}
