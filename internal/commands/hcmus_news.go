package commands

import (
	"fit-news-discord-bot/internal/repository"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HcmusNewsCommand() Command {
	return Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "hcmus_news",
			Description: "Lấy tin tức mới nhất từ HCMUS",
		},
		Handler: func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
			news, err := repository.GetAllNews("FITHCMUS")
			if err != nil || len(news) == 0 {
				s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Lỗi lấy tin tức hoặc không có tin tức nào.",
					},
				})
				return
			}

			response := fmt.Sprintf("📰 | **%s**\n\n%s\n\n", news[0].Title, news[0].URL)

			s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
		},
	}
}
