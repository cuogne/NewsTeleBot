package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HelloCommand() Command {
	return Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "hello",
			Description: "Says hello to the world",
		},
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello World!",
				},
			})
		},
	}
}
