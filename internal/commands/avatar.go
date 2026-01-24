package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func GetAvatarCommand() Command {
	return Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "avatar",
			Description: "Lấy avatar của user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "Người dùng để lấy avatar",
					Required:    false,
				},
			},
		},
		Handler: func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
			var user *discordgo.User

			if len(interaction.ApplicationCommandData().Options) > 0 {
				userOption := interaction.ApplicationCommandData().Options[0]
				user = userOption.UserValue(s)
			}

			if user == nil {
				user = interaction.Member.User // lay avatar cua nguoi gui
			}

			avatarURL := user.AvatarURL("1024")

			embed := &discordgo.MessageEmbed{
				Color: 0x0099ff,
				Title: interaction.Member.User.Username + " muốn xem ảnh của " + user.Username,
				Image: &discordgo.MessageEmbedImage{
					URL: avatarURL,
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Requested by " + interaction.Member.User.Username,
					IconURL: interaction.Member.User.AvatarURL(""),
				},
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			}

			s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
		},
	}
}
