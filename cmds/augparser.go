package cmds

import (
	"fmt"
	"strings"
)

// AugmentParser Parses input text for augments.
type AugmentParser struct {
	// registry to query for augment definitions
	registry Registry
}

func NewAugmentParser(registry Registry) *AugmentParser {
	return &AugmentParser{registry}
}

// Parse augments
func (p AugmentParser) Parse(tokens []string, cmdReq *CmdReq) error {
	// Check tokens not empty
	if len(tokens) == 0 {
		return nil
	}

	// Check if first token is augment
	token := CleanToken(tokens[0])

	// Resplit token just in case punctuation was gluing words together
	newTokens := strings.Fields(token)

	// If there are new tokens
	if len(newTokens) > 1 {
		// Pop current token
		tokens = tokens[1:]

		// Replace just popped token with space split version
		tokens = append(newTokens, tokens...)

		// Set new token
		token = tokens[0]
	}

	if augment := p.registry.Augment(token); augment != nil {
		// Check that this augment hasn't been provided before under another
		// value
		if val, ok := cmdReq.Augments[augment.Name]; ok {
			return fmt.Errorf("augment \"%s\" already provided with value \"%s\", duplicate value: \"%s\"", augment.Name, val, token)
		}

		// Add augment to cmdReq
		cmdReq.Augments[augment.Name] = token

		// Remove token we just processed
		tokens = tokens[1:] // pop

		// Recurse
		return p.Parse(tokens, cmdReq)
	} else if p.registry.Cmd(token) == nil {
		// If token is not command or augment, ignore
		tokens = tokens[1:] // pop

		// Recurse
		return p.Parse(tokens, cmdReq)
	} else {
		// If not augment, but not command, stop looking for augments
		return nil
	}
}
