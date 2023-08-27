package component_helper

import "github.com/bwmarrin/discordgo"

func CreateSelectMenu(placeholder string, customId string, options []discordgo.SelectMenuOption) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.SelectMenu{
					MenuType:    discordgo.StringSelectMenu,
					CustomID:    customId,
					Placeholder: placeholder,
					Options:     options,
				},
			},
		},
	}
}
