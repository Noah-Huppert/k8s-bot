package chat

type Augment string

const (
	// Thread forces the bot to reply in a thread
	Thread Augment = "thread"

	// Channel forces the bot to reply in a channel
	Channel Augment = "channel"

	// Pm forces the bot to reply in a private message
	Pm Augment = "pm"
)
