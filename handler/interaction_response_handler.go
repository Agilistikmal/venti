package handler

import (
	"github.com/agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func RespondInteractionSuccess(bot *discordgo.Session, interaction *discordgo.Interaction, message string) {
	bot.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed_helper.SuccessWithMessage(message),
			},
			Flags: 1 << 6,
		},
	})
}

func RespondInteractionError(bot *discordgo.Session, interaction *discordgo.Interaction, message string) {
	bot.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed_helper.ErrorWithMessage(message),
			},
			Flags: 1 << 6,
		},
	})
}
