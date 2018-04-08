package chat

// CmdReq holds information about command requests received by the bot. Values
// populated by Parse()
type CmdReq struct {
	// RawText holds the unprocessed message text
	RawText string

	// Sender holds the ID of the user who sent the message
	Sender string

	// Channel is the slack channel the message was received in
	Channel string

	// Relevant indicates if the bot should take action based on this message
	Relevant bool

	// Augments holds the augments present in the command request
	Augments map[string]string

	// Cmd holds the name of the command specified in the request
	Cmd string

	// Args holds the arguments provided in the command request. Keys are arg
	// names.
	Args map[string]string
}

// NewCmdReq initialized a CmdReq with blank values
func NewCmdReq() *CmdReq {
	var cmdReq CmdReq

	cmdReq.Augments = make(map[string]string)
	cmdReq.Args = make(map[string]string)

	return &cmdReq
}
