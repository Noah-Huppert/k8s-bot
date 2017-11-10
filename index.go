package main

import (
	"./bot"
	"./config"
	"fmt"
)

func main() {
	// Load config
	config := config.LoadConfig()
	if err != nil {
		fmt.Printf("error loading config: %s\n", err.Error())
		return
	}

	// Start Bot
	kube := bot.NewBot(config)
	if err := kube.Start(); err != nil {
		fmt.Printf("error starting bot: %s\n", err.Error())
		return
	}

	fmt.Println("started bot")

	// Stop Bot
	fmt.Println("stopping bot")

	if err := kube.Stop(); err != nil {
		fmt.Printf("error stopping bot: %s\n", err.Error())
		return
	}

	fmt.Println("stopped bot")
}
