package chat

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

// AllParser runs all the defined parsers in the correct order.
type AllParser struct {
	// registry is used to save command and augmentation definitions
	registry Registry

	// slackAPI is the client used to communicate with Slack
	slackAPI *slack.Client

	// slackRTM is the real time messaging client used to communicate with Slack
	slackRTM *slack.RTM
}

// NewAllParser creates and returns a AllParser instance
func NewAllParser(registry Registry, slackAPI *slack.Client, slackRTM *slack.RTM) *AllParser {
	p := &AllParser{}

	p.registry = registry

	p.slackAPI = slackAPI
	p.slackRTM = slackRTM

	return p
}

// ParseText interprets received message text and returns a CmdReq struct
// which indicates the command the user sent.
// An error is returned if one occurs, nil on success.
func (p AllParser) Parse(msg slack.Msg) (*CmdReq, error) {
	// Make empty command request
	cmdReq := NewCmdReq()

	// Check msg not empty
	if len(msg.Text) == 0 {
		return nil, errors.New("message text is empty")
	}

	// Parse
	// Channel
	if len(msg.Channel) == 0 {
		return nil, errors.New("message channel is empty")
	}
	cmdReq.Channel = msg.Channel

	// Split by spaces
	tokens := strings.Fields(msg.Text)

	// Relevance
	relevanceParser := NewRelevanceParser(p.slackAPI, p.slackRTM)
	if err := relevanceParser.Parse(tokens, cmdReq); err != nil {
		return nil, fmt.Errorf("error determining message relevance: %s", err.Error())
	}

	// Augments
	augParser := NewAugmentParser(p.registry)
	if err := augParser.Parse(tokens, cmdReq); err != nil {
		return nil, fmt.Errorf("error parsing augments: %s", err.Error())
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
