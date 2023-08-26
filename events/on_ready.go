package events

import (
	"github.com/agilistikmal/venti/handler"
	"github.com/bwmarrin/discordgo"
	"log"
)

func OnReady(bot *discordgo.Session, event *discordgo.Ready) {
	err := bot.UpdateWatchStatus(0, "gprestore.net")
	handler.HandleError(err)
	log.Printf("Discord bot (%s) is online!", bot.State.User.String())
}
