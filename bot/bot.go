package bot

import (
	"../config"
	"context"
)

// Bot acts as a chat bot based interface to the Kubernetes API. Leveraging the
// Slack API to send and receive messages.
type Bot struct {
	// Config holds the setting values provided by the user. Used to tweak the
	// bot's behavior.
	Config config.Config

	// ctx is used to stop the Bot's execution. The Bot will only run while the
	// context is not expired.
	ctx context.Context

	// Name is the Slack username the Bot will respond to
	Name string
}

// NewBot creates a new Bot instance from the parameters specified in the
// Config object
func NewBot(config config.Config) Bot {
	return Bot{Name: config.bot.name}
}

// Start begins the process of receiving and responding to user messages.
// Returns nil if Bot starts successfully, or the error that occurred otherwise.
func (b Bot) Start() error {
	return nil
}

// Stop ends the process of receiving and responding to user messages. Returns
// an error if one occurs. Nil on success.
func (b Bot) Stop() error {
	return nil
}
