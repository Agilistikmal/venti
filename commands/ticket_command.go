package commands

import (
	"github.com/agilistikmal/venti/config"
	"github.com/agilistikmal/venti/helper"
	"github.com/agilistikmal/venti/helper/component_helper"
	"github.com/agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func TicketCommand(bot *discordgo.Session, message *discordgo.MessageCreate) {
	command, args := helper.GetCommand(message.Content)
	if command == "ticket" {
		if args[0] == "panel" {
			if helper.Contains(message.Member.Roles, config.StaffRoleId) == false {
				bot.ChannelMessageSendEmbedReply(message.ChannelID, embed_helper.NoAccessError(), message.Reference())
				return
			}
			channelId := helper.ChannelMentionToChannelId(args[1])
			embed := embed_helper.CustomWithTimestamp(
				"",
				"## Ticket\n### Pusat Informasi dan Bantuan\nJika ada yang perlu ditanyakan silahkan klik tombol **Open Ticket** dibawah ini.\n\n**BUKAN UNTUK PEMBELIAN**, pembelian melalui channel <#1039148098808725624>",
				embed_helper.BLUE,
			)
			components := component_helper.CreateButton("Open Ticket", discordgo.PrimaryButton, "open-ticket")
			_, err := bot.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
				Embed:      embed,
				Components: components,
			})
			if err != nil {
				bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError("!ticket panel <channel>"))
			}
		}
	}
}
