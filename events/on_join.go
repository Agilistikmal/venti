package events

import (
	"fmt"
	"github.com/Agilistikmal/venti/config"
	"github.com/Agilistikmal/venti/handler"
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/bwmarrin/discordgo"
)

func OnJoin(bot *discordgo.Session, event *discordgo.GuildMemberAdd) {
	bot.ChannelMessageSendComplex(config.WelcomeChannelId, &discordgo.MessageSend{
		Content: event.Member.Mention(),
		Embed: embed_helper.CustomWithFields(
			event.User,
			"",
			fmt.Sprintf("Selamat datang di server discord **GPrestore** %s.\n\n- <#1039148098808725624> untuk pembelian\n- <#1041685950952124446> untuk melihat rating\n- <#1144821480002170960> untuk menghubungi kami\n\n_Mohon menjaga perkataan demi kenyamanan bersama._", event.User.Mention()),
			[]*discordgo.MessageEmbedField{
				embed_helper.CreateField("Shopee", "[Lihat Shopee](https://shopee.co.id/gprestore)", true),
				embed_helper.CreateField("Whatsapp", "[Lihat Whatsapp](https://wa.me/6285888881550)", true),
				embed_helper.CreateField("Instagram (tidak aktif)", "[Lihat Instagram](https://instagram.com/gprestoreid)", true),
			},
			embed_helper.WHITE,
		),
	})
	err := bot.GuildMemberRoleAdd(event.GuildID, event.User.ID, config.UserRoleId)
	handler.HandleError(err)
}
