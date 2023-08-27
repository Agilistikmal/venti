package commands

import (
	"github.com/Agilistikmal/venti/config"
	"github.com/Agilistikmal/venti/helper"
	"github.com/Agilistikmal/venti/helper/component_helper"
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func FAQCommand(bot *discordgo.Session, message *discordgo.MessageCreate) {
	command, args := helper.GetCommand(message.Content)
	if command != "faq" {
		return
	}
	if len(args) < 1 {
		bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError(config.Prefix+"faq panel <channel>"))
		return
	}
	if args[0] == "panel" {
		if helper.Contains(message.Member.Roles, config.StaffRoleId) == false {
			bot.ChannelMessageSendEmbedReply(message.ChannelID, embed_helper.NoAccessError(), message.Reference())
			return
		}
		if len(args) < 2 {
			bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError(config.Prefix+"faq panel <channel>"))
			return
		}
		channelId := helper.ChannelMentionToChannelId(args[1])
		embed := embed_helper.CustomWithTimestamp(
			"FAQ (Frequently Asked Question)",
			"Berikut adalah pertanyaan yang sering ditanyakan. Anda bisa melihat jawaban dari daftar pertanyaan dibawah ini.",
			embed_helper.WHITE,
		)
		components := component_helper.CreateSelectMenu("Pilih pertanyaan disini", "faq-list", config.FAQList)
		_, err := bot.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
			Embed:      embed,
			Components: components,
		})
		if err != nil {
			bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.ErrorWithMessage("Channel tidak ditemukan."))
		}
	}
}
