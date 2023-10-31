package product_event

import (
	"github.com/Agilistikmal/venti/config"
	"github.com/Agilistikmal/venti/handler"
	"github.com/Agilistikmal/venti/helper/component_helper"
	"github.com/Agilistikmal/venti/service"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func PaymentConfirmButtonClick(bot *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionMessageComponent {
		return
	}
	if interaction.MessageComponentData().CustomID != "confirm-payment" {
		return
	}
	embed := interaction.Message.Embeds[0]
	externalId := strings.Split(embed.Title, " ")[1]
	payment, err := service.GetPayment(externalId)
	if err != nil {
		handler.RespondInteractionError(bot, interaction.Interaction, err.Error())
		return
	}

	embed.Description = "Email: " + strings.Split(embed.Description, "\n")[0]
	embed.Image = nil
	embed.Fields = []*discordgo.MessageEmbedField{
		embed.Fields[0],
		embed.Fields[1],
		embed.Fields[2],
		embed.Fields[3],
	}

	if payment[0].Status == "SUCCESS" {
		bot.ChannelMessageSendComplex(config.PurchaseLogChannelId, &discordgo.MessageSend{
			Components: component_helper.CreateButton("Send", discordgo.SecondaryButton, "send-product"),
			Embed:      embed,
		})
		bot.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Components: []discordgo.MessageComponent{},
			Embeds:     interaction.Interaction.Message.Embeds,
			ID:         interaction.Interaction.Message.ID,
			Channel:    interaction.Interaction.ChannelID,
		})
		handler.RespondInteractionSuccess(bot, interaction.Interaction, "Berhasil! Pesanan telah kami terima. Mohon tunggu admin akan memproses pesanan anda dan akan dikirim melalui DM bot ini. Terimakasih", false)
	} else {
		handler.RespondInteractionError(bot, interaction.Interaction, "Status belum dibayar")
	}
}
