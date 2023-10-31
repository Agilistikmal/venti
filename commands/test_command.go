package commands

import (
	"github.com/Agilistikmal/venti/helper"
	"github.com/Agilistikmal/venti/service"
	"github.com/bwmarrin/discordgo"
	"log"
)

func TestCMD(bot *discordgo.Session, message *discordgo.MessageCreate) {
	command, _ := helper.GetCommand(message.Content)
	if command != "test" {
		return
	}
	products := service.FindAllProduct()
	bot.ChannelMessageSend(message.ChannelID, "Success")
	log.Print(products[0])
}
