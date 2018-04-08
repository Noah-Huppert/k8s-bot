package chat

// Parser interprets text sent by the user and returns a CmdReq representing
// the command the user requested to run
type Parser interface {
	// Parse interprets the provided tokens, and saves their meaning in the
	// CmdReq pointer. Tokens that are parsed should be removed from the
	// tokens slice. An error will be returned if one occurs, nil on success.
	Parse(tokens []string, cmdReq *CmdReq) error
}
