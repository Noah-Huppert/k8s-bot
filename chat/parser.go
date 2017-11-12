package chat

import (
	"errors"
	"fmt"
	"strings"
)

// Parser interprets text sent by the user and returns a CmdReq representing
// the command the user requested to run
type Parser interface {
	// Parse interprets the provided tokens, and saves their meaning in the
	// CmdReq pointer. Tokens that are parsed should be removed from the
	// tokens slice. An error will be returned if one occurs, nil on success.
	Parse(tokens []string, cmdReq *CmdReq) error
}

// ParseText interprets received message text and returns a CmdReq struct
// which indicates the command the user sent.
// An error is returned if one occurs, nil on success.
func ParseText(txt string, registry Registry) (CmdReq, error) {
	cmdReq := EmptyCmdReq()

	// Check not empty
	if len(txt) == 0 {
		return cmdReq, errors.New("error parsing text, was empty")
	}

	// Split by spaces
	tokens := strings.Fields(txt)

	// Augments
	p := NewAugmentParser(registry)
	if err := p.Parse(tokens, &cmdReq); err != nil {
		return cmdReq, fmt.Errorf("error parsing augments: %s", err.Error())
	}

	// Done
	return cmdReq, nil
}

func CleanToken(txt string) string {
	ignoredSymbols := []string{
		",",
		".",
	}

	// remove ignored symbols
	for _, sym := range ignoredSymbols {
		txt = strings.Replace(txt, sym, " ", -1)
	}

	return txt
}
