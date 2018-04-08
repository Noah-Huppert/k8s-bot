package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/kube-bot/bot"
	"github.com/Noah-Huppert/kube-bot/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// Metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":8080", nil); err != nil {
			logger.Printf("error serving metrics: %s", err.Error())
		}
	}()

	// Run Bot
	logger.Println("starting bot")

	if err := kube.Run(); err != nil {
		logger.Printf("error running bot: %s\n", err.Error())
		return
	}
}
