package commands

import (
	"github.com/bwmarrin/discordgo"
)

func PingCommand() Command {
	return Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Responds with Pong!",
		},
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
	}
}
