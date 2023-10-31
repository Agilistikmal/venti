package product_event

import (
	"github.com/Agilistikmal/venti/handler"
	"github.com/Agilistikmal/venti/helper/component_helper"
	"github.com/Agilistikmal/venti/service"
	"github.com/bwmarrin/discordgo"
)

func PurchaseButtonClick(bot *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionMessageComponent {
		return
	}
	if interaction.MessageComponentData().CustomID != "purchase" {
		return
	}
	productId := interaction.Message.Content
	product, _ := service.FindProductById(productId)
	if product.Stock == 0 {
		handler.RespondInteractionError(bot, interaction.Interaction, "Stock kosong.")
		return
	}
	err := bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			Components: component_helper.CreateModalPurchase(),
			Flags:      1 << 6,
			CustomID:   "purchase-confirm|" + productId,
			Title:      product.Name,
		},
	})
	if err != nil {
		handler.RespondInteractionError(bot, interaction.Interaction, "Internal server error (purchase modal)")
	}
}
