package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"fit-news-discord-bot/internal/config"
)

func main() {
	// load config discord
	discordCfg, err := config.LoadDiscordConfig()
	if err != nil {
		log.Fatalf("Error loading discord config: %v", err)
	}

	// initiate Discord session
	session, err := discordCfg.Connect()
	if err != nil {
		log.Fatalf("Error connecting to Discord: %v", err)
	}

	// setup bot commands
	registry := config.NewRegistry()
	config.SetupCommands(registry)

	// add interaction handler
	session.AddHandler(registry.HandleInteraction)

	// start the bot
	if err := session.Open(); err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer session.Close()

	log.Println("Bot is running...")

	// register commands
	if err := config.DeployCommands(session, registry); err != nil {
		log.Fatalf("Cannot register commands: %v", err)
	}

	// 7. Wait for Interrupt Signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Shutting down...")
}
