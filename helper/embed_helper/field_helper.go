package embed_helper

import "github.com/bwmarrin/discordgo"

func CreateField(name string, value string, inline bool) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
}
