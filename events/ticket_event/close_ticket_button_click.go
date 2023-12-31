package ticket_event

import "github.com/bwmarrin/discordgo"

func CloseTicketButtonClick(bot *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionMessageComponent {
		return
	}
	if interaction.MessageComponentData().CustomID != "close-ticket" {
		return
	}
	bot.ChannelDelete(interaction.ChannelID)
}
