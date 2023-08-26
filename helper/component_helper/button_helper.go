package component_helper

import "github.com/bwmarrin/discordgo"

func CreateButton(label string, style discordgo.ButtonStyle, customId string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    label,
					Style:    style,
					CustomID: customId,
				},
			},
		},
	}
}

func CreateButtonLink(label string, style int, customId string, link string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    label,
					Style:    discordgo.ButtonStyle(style),
					CustomID: customId,
					URL:      link,
				},
			},
		},
	}
}
