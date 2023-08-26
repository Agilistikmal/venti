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
