package embed_helper

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func CustomWithTimestamp(title string, description string, color ...int) *discordgo.MessageEmbed {
	embedColor := TRANSPARENT
	if len(color) >= 1 {
		embedColor = color[0]
	}
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       embedColor,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}

func CustomWithFields(author *discordgo.User, title string, description string, fields []*discordgo.MessageEmbedField, color ...int) *discordgo.MessageEmbed {
	embedColor := TRANSPARENT
	if len(color) >= 1 {
		embedColor = color[0]
	}
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    author.Username,
			IconURL: author.AvatarURL("64"),
		},
		Title:       title,
		Description: description,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       embedColor,
		Fields:      fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}
