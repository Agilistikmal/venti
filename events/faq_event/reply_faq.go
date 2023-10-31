package faq_event

import (
	"github.com/Agilistikmal/venti/config"
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func ReplyFAQ(bot *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionMessageComponent {
		return
	}
	if interaction.MessageComponentData().CustomID != "faq-list" {
		return
	}
	faqValue := interaction.MessageComponentData().Values[0]
	faqResponses := config.FAQResponse
	for _, faq := range faqResponses {
		if faq.Value == faqValue {
			bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						embed_helper.CustomWithTimestamp(faq.Value, faq.Response),
					},
					Flags: 1 << 6,
				},
			})
		}
	}
}
