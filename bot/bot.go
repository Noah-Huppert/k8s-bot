package bot

import (
	"context"
	"log"
	"os"

	"github.com/Noah-Huppert/kube-bot/config"
	"github.com/nlopes/slack"
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

	// ctxCancelFn is the cancel function for ctx
	ctxCancelFn context.CancelFunc

	// Name is the Slack username the Bot will respond to
	Name string

	// logger is used to record debug messages
	logger *log.Logger

	// slackLogger is the logger used by the Slack API client
	slackLogger *log.Logger

	// slackAPI is the authenticated Slack API client used to interact with
	// Slack
	slackAPI *slack.Client

	// slackRTM is the Slack API real time messaging client used to receive and
	// respond to Slack API events
	slackRTM *slack.RTM
}

// NewBot creates a new Bot instance from the parameters specified in the
// Config object
func NewBot(ctx context.Context, cfg config.Config) Bot {
	// Loggers
	logger := log.New(os.Stdout, "bot: ", 0)
	slackLogger := log.New(os.Stdout, "bot: slack api: ", 0)

	// Context
	ctx, ctxCancelFn := context.WithCancel(ctx)

	// Make
	return Bot{
		Name:        cfg.Bot.Name,
		logger:      logger,
		slackLogger: slackLogger,
		ctx:         ctx,
		ctxCancelFn: ctxCancelFn,
		Config:      cfg,
	}
}

// Run begins the process of receiving and responding to user messages. This
// process will continue until Stop() is called.
//
// Returns the error that stopped execution. If Run() was stopped by a context
// either the context.Canceled or context.DeadlineExceeded error will be
// returned.
func (b Bot) Run() error {
	b.logger.Println("running bot")

	// Init Slack Lib
	b.slackAPI = slack.New(b.Config.SlackAPIToken)
	slack.SetLogger(b.slackLogger)

	// Connect
	b.slackRTM = b.slackAPI.NewRTM()
	go b.slackRTM.ManageConnection()

	// Receive
	go b.handleEvents(b.slackRTM.IncomingEvents)

	select {
	case <-b.ctx.Done():
		b.logger.Printf("received shutdown request: %s", b.ctx.Err().Error())
		return b.ctx.Err()
	}

	return nil
}

// handleEvents receives Slack events via the provided channel and processes
// them accordingly. Returns the error that stopped execution.
func (b Bot) handleEvents(in <-chan slack.RTMEvent) error {
	b.logger.Println("starting to receive Slack events")

	for {
		select {
		case <-b.ctx.Done():
			// Ctx has expired
			return b.ctx.Err()
		case msg := <-in:
			// Received Slack API event
			switch event := msg.Data.(type) {
			case *slack.MessageEvent:
				if err := b.handleMessage(event); err != nil {
					b.logger.Printf("error handling message: %s\n", err.Error())
				}
			case *slack.InvalidAuthEvent:
				b.logger.Println("invalid credentials")
			}
		}
	}

	return nil
}

// handleMessage performs the appropriate actions for the provided message event.
// Returns an error on failure, nil on success.
func (b Bot) handleMessage(event *slack.MessageEvent) error {
	msg := event.Msg

	// Log
	b.logger.Printf("received message: %s\n", msg.Text)

	// Echo
	out := b.slackRTM.NewOutgoingMessage(msg.Text, event.Channel)
	b.slackRTM.SendMessage(out)

	return nil
}

// Stop ends the process of receiving and responding to user messages. This
// will cause the Run() method to exit and return a context.Canceled error.
func (b Bot) Stop() {
	b.ctxCancelFn()
}
