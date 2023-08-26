package embed_helper

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func NoAccessError() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: "No Access Error! Anda tidak memiliki akses untuk ini.",
		Color:       RED,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}

func InvalidUsageError(usage string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: fmt.Sprintf("Invalid Usage! Gunakan `%s`", usage),
		Color:       RED,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}

func InternalServerError() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: "Internal Server Error! Silahkan dicoba kembali.",
		Color:       RED,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}

func ErrorWithMessage(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: message,
		Color:       RED,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Venti v1 by safatanc.com",
		},
	}
}
