package embed_helper

import "github.com/bwmarrin/discordgo"

func SuccessWithMessage(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Success",
		Description: message,
		Color:       RED,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}
