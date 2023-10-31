package main

import (
	"github.com/Agilistikmal/venti/commands"
	"github.com/Agilistikmal/venti/database"
	"github.com/Agilistikmal/venti/events"
	"github.com/Agilistikmal/venti/events/faq_event"
	"github.com/Agilistikmal/venti/events/product_event"
	"github.com/Agilistikmal/venti/events/ticket_event"
	"github.com/Agilistikmal/venti/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load()
	handler.HandleError(err)

	database.CreateConnection()
	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")
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
	bot.AddHandler(events.OnJoin)

	// Commands
	bot.AddHandler(commands.TicketCommand)
	bot.AddHandler(commands.FAQCommand)
	bot.AddHandler(commands.PurchaseCommand)

	// Component Events
	bot.AddHandler(ticket_event.OpenTicketButtonClick)
	bot.AddHandler(ticket_event.CloseTicketButtonClick)
	bot.AddHandler(faq_event.ReplyFAQ)
	bot.AddHandler(product_event.PurchaseButtonClick)
	bot.AddHandler(product_event.PurchaseModalSubmit)
	bot.AddHandler(product_event.PaymentConfirmButtonClick)
}
