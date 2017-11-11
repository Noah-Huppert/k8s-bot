package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/kube-bot/bot"
	"github.com/Noah-Huppert/kube-bot/config"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "main: ", 0)

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Printf("error loading config: %s\n", err.Error())
		return
	}

	// Create Bot
	ctx := context.Background()
	kube, err := bot.NewBot(ctx, cfg)
	if err != nil {
		logger.Printf("error creating bot: %s\n", err.Error())
	}

	// Handle system Signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for range sigChan {
			logger.Println("received SIGINT")

			kube.Stop()
		}
	}()

	// Run Bot
	logger.Println("starting bot")

	if err := kube.Run(); err != nil {
		logger.Printf("error running bot: %s\n", err.Error())
		return
	}
}
