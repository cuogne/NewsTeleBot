package config

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type DiscordConfig struct {
	Token string
}

func LoadDiscordConfig() (*DiscordConfig, error) {
	_ = godotenv.Load()

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("DISCORD_BOT_TOKEN is required")
	}

	return &DiscordConfig{
		Token: token,
	}, nil
}

// establishes a connection to discord and returns the session
func (c *DiscordConfig) Connect() (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsGuildMessages

	return session, nil
}
