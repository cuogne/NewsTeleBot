package config

import (
	"fit-news-discord-bot/internal/commands"

	"github.com/bwmarrin/discordgo"
)

type Registry struct {
	Commands map[string]commands.Command
}

func NewRegistry() *Registry {
	return &Registry{
		Commands: make(map[string]commands.Command),
	}
}

func (r *Registry) Register(name string, cmd commands.Command) {
	r.Commands[name] = cmd
}

func (r *Registry) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	cmd, ok := r.Commands[i.ApplicationCommandData().Name]
	if !ok {
		return
	}

	cmd.Handler(s, i)
}

func DeployCommands(s *discordgo.Session, r *Registry) error {
	var cmds []*discordgo.ApplicationCommand
	for _, cmd := range r.Commands {
		cmds = append(cmds, cmd.Definition)
	}

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", cmds)
	return err
}
