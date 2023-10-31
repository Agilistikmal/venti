package component_helper

import "github.com/bwmarrin/discordgo"

func CreateModalPurchase() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID:    "email",
					Label:       "Email",
					Style:       discordgo.TextInputShort,
					Placeholder: "Email untuk pengiriman",
					Required:    true,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID:    "quantity",
					Label:       "Jumlah Pembelian",
					Style:       discordgo.TextInputShort,
					Placeholder: "1",
					Value:       "1",
					Required:    true,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID:    "notes",
					Label:       "Catatan (opsional)",
					Style:       discordgo.TextInputParagraph,
					Placeholder: "Kosongkan jika tidak ada",
					Required:    false,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID:    "voucher",
					Label:       "Voucher (opsional)",
					Style:       discordgo.TextInputShort,
					Placeholder: "Kosongkan jika tidak ada",
					Required:    false,
				},
			},
		},
	}
}
