package commands

import (
	"github.com/Agilistikmal/venti/config"
	"github.com/Agilistikmal/venti/helper"
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/Agilistikmal/venti/service"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func PurchaseCommand(bot *discordgo.Session, message *discordgo.MessageCreate) {
	command, args := helper.GetCommand(message.Content)
	if command != "purchase" {
		return
	}
	if len(args) < 1 {
		bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError(config.Prefix+"purchase panel <channel>"))
		return
	}
	if args[0] == "panel" {
		if helper.ContainOnList(message.Member.Roles, config.StaffRoleId) == false {
			bot.ChannelMessageSendEmbedReply(message.ChannelID, embed_helper.NoAccessError(), message.Reference())
			return
		}
		if len(args) < 2 {
			bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.InvalidUsageError(config.Prefix+"purchase panel <channel>"))
			return
		}

		channelId := helper.ChannelMentionToChannelId(args[1])
		products := service.FindAllProduct()

		for _, product := range products {
			embed := embed_helper.CustomWithFields(
				&discordgo.User{},
				product.Name,
				"```"+product.Description+"```"+"```💰 Rp"+strconv.Itoa(product.Price)+"```",
				[]*discordgo.MessageEmbedField{
					embed_helper.CreateField("Stock", "` 📦 "+strconv.Itoa(product.Stock)+" `", true),
					embed_helper.CreateField("Stars", "` ⭐ "+strconv.Itoa(product.Stars)+" `", true),
					embed_helper.CreateField("Sold", "` Terjual "+strconv.Itoa(product.Stars)+"x `", true),
				},
				embed_helper.WHITE,
			)
			components := []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Beli",
							Style:    discordgo.SecondaryButton,
							CustomID: "purchase",
						},
						discordgo.Button{
							Label:    "Beri ulasan",
							Style:    discordgo.SecondaryButton,
							CustomID: "rate",
						},
					},
				},
			}
			_, err := bot.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
				Embed:      embed,
				Components: components,
				Content:    product.ID,
			})
			if err != nil {
				bot.ChannelMessageSendEmbed(message.ChannelID, embed_helper.ErrorWithMessage("Channel tidak ditemukan."))
			}
		}
	}
}
