package chat

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

// RelevanceParser determines if a message is directed towards the bot. A message
// is relevant if:
//      - @ mentions the bot
//      - pm's the bot
type RelevanceParser struct {
	// slackRTM is the real time messaging client used to communicate with Slack
	slackAPI *slack.Client
	slackRTM *slack.RTM

	// logger is used to print debug information about message relevance
	logger *log.Logger
}

// Parse implements the Parser interface for the RelevanceParser
func (p RelevanceParser) Parse(tokens []string, cmdReq *CmdReq) error {
	// Get bot info
	info := p.slackRTM.GetInfo()

	// Print
	p.logger.Printf("message: channel => %s, bot: id => %s, name => %s, ims => %+v",
		cmdReq.Channel, info.User.ID, info.User.Name, info.IMs)

	return nil
}

// NewRelevanceParser creates a new RelevanceParser instance
func NewRelevanceParser(slackAPI *slack.Client, slackRTM *slack.RTM) *RelevanceParser {
	p := &RelevanceParser{}

	p.slackAPI = slackAPI
	p.slackRTM = slackRTM

	p.logger = log.New(os.Stdout, "bot: relevance: ", 0)

	return p
}
