package component_event

import (
	"fmt"
	"github.com/agilistikmal/venti/config"
	"github.com/agilistikmal/venti/handler"
	"github.com/agilistikmal/venti/helper/component_helper"
	"github.com/agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func OpenTicketButtonClick(bot *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.MessageComponentData().CustomID != "open-ticket" {
		return
	}
	channels, err := bot.GuildChannels(interaction.GuildID)
	var channelExist string
	for _, channel := range channels {
		if channel.Name == "ticket-"+interaction.Member.User.Username {
			channelExist = channel.Mention()
		}
	}
	if channelExist != "" {
		handler.RespondInteractionError(bot, interaction.Interaction, fmt.Sprintf("Anda masih memiliki ticket yang terbuka. Silahkan ke ticket anda %s", channelExist))
		return
	}
	ticketChannel, err := bot.GuildChannelCreateComplex(interaction.GuildID, discordgo.GuildChannelCreateData{
		Name:     "ticket-" + interaction.Member.User.Username,
		Position: 0,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    interaction.Member.User.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: 1 << 10,
			},
			{
				ID:   "967621643058942024",
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: 1 << 10,
			},
		},
		ParentID: config.TicketCategoryId,
	})
	if err != nil {
		handler.RespondInteractionError(bot, interaction.Interaction, err.Error())
	}
	_, err = bot.ChannelMessageSendComplex(ticketChannel.ID,
		&discordgo.MessageSend{
			Content: interaction.Member.Mention() + "<@&970645243546435644>",
			Embed: embed_helper.CustomWithTimestamp(
				"Ticket",
				"Silahkan ajukan pertanyaan dibawah ini. Admin akan segera menjawab pertanyaan anda.\n*Transcript ticket ini tidak akan disimpan. Sebelum di close sebaiknya menyimpan Informasi penting agar tidak hilang.",
			),
			Components: component_helper.CreateButton("Close Ticket", discordgo.DangerButton, "close-ticket"),
		},
	)
	if err != nil {
		handler.RespondInteractionError(bot, interaction.Interaction, err.Error())
	}
	bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed_helper.SuccessWithMessage(
					fmt.Sprintf("Berhasil membuat ticket, silahkan ke %s untuk melihat ticket anda.", ticketChannel.Mention()),
				),
			},
			Flags: 1 << 6,
		},
	})
}
