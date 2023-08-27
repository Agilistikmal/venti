package main

import (
	"github.com/Agilistikmal/venti/commands"
	"github.com/Agilistikmal/venti/events"
	"github.com/Agilistikmal/venti/events/faq_event"
	"github.com/Agilistikmal/venti/events/ticket_event"
	"github.com/Agilistikmal/venti/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load()
	handler.HandleError(err)

	token := os.Getenv("BOT_TOKEN")
	bot, err := discordgo.New("Bot " + token)
	handler.HandleError(err)

	bot.Identify.Intents = discordgo.IntentsAll
	RegisterEvents(bot)

	err = bot.Open()
	handler.HandleError(err)
	defer bot.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func RegisterEvents(bot *discordgo.Session) {
	// Events
	bot.AddHandler(events.OnReady)

	// Commands
	bot.AddHandler(commands.TicketCommand)
	bot.AddHandler(commands.FAQCommand)

	// Component Events
	bot.AddHandler(ticket_event.OpenTicketButtonClick)
	bot.AddHandler(ticket_event.CloseTicketButtonClick)
	bot.AddHandler(faq_event.ReplyFAQ)
}
