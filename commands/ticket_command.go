package commands

import (
	"github.com/Agilistikmal/venti/config"
	"github.com/Agilistikmal/venti/helper"
	"github.com/Agilistikmal/venti/helper/component_helper"
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func TicketCommand(bot *discordgo.Session, message *discordgo.MessageCreate) {
	command, args := helper.GetCommand(message.Content)
	if command != "ticket" {
		return
	}
	if len(args) < 1 {
		bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError(config.Prefix+"ticket panel <channel>"))
		return
	}
	if args[0] == "panel" {
		if helper.ContainOnList(message.Member.Roles, config.StaffRoleId) == false {
			bot.ChannelMessageSendEmbedReply(message.ChannelID, embed_helper.NoAccessError(), message.Reference())
			return
		}
		if len(args) < 2 {
			bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError(config.Prefix+"ticket panel <channel>"))
			return
		}
		channelId := helper.ChannelMentionToChannelId(args[1])
		embed := embed_helper.CustomWithTimestamp(
			"Pusat Informasi dan Bantuan",
			"Jika ada yang perlu ditanyakan silahkan klik tombol **Open Ticket** dibawah ini.\n\n*__Bukan untuk pembelian__, pembelian melalui channel <#1039148098808725624>",
			embed_helper.WHITE,
		)
		components := component_helper.CreateButton("Open Ticket", discordgo.SecondaryButton, "open-ticket")
		_, err := bot.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
			Embed:      embed,
			Components: components,
		})
		if err != nil {
			bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.ErrorWithMessage("Channel tidak ditemukan."))
		}
	}
}
