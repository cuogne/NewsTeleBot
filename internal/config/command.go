package config

import "fit-news-discord-bot/internal/commands"

func SetupCommands(registry *Registry) {
	registry.Register("ping", commands.PingCommand())
	registry.Register("avatar", commands.GetAvatarCommand())
}
