package handler

import (
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func RespondInteractionSuccess(bot *discordgo.Session, interaction *discordgo.Interaction, message string, ephemeral ...bool) {
	messageFlags := discordgo.MessageFlagsEphemeral
	if len(ephemeral) >= 1 {
		if ephemeral[0] == false {
			messageFlags = discordgo.MessageFlagsIsCrossPosted
		} else {
			messageFlags = discordgo.MessageFlagsEphemeral
		}
	}
	bot.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed_helper.SuccessWithMessage(message),
			},
			Flags: messageFlags,
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
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
