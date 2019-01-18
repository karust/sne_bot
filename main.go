package main

import (
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func retErr(err string, descr ...error) {
	if len(descr) > 0 {
		print("["+err+"]: ", descr, "\n")
	} else {
		print("[" + err + "]\n")
	}
	os.Exit(1)
}

var (
	// Bot ... Telegram API bot instanse
	Bot tgbotapi.BotAPI
)

func main() {
	// export TELEGRAM_APITOKEN=...
	//Token = os.Getenv("TELEGRAM_APITOKEN")
	if Token == "" {
		retErr("Error accessing token")
	}
	Bot, err := tgbotapi.NewBotAPI(string(Token))
	if err != nil {
		retErr("Can't load API token", err)
	}

	// Display every request and response.
	Bot.Debug = true

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	// conf - Struct with configs for receiving updates
	conf := tgbotapi.NewUpdate(0)
	conf.Timeout = 60

	updates, err := Bot.GetUpdatesChan(conf)

	// Wait for old updates and clear them
	time.Sleep(time.Millisecond * 100)
	updates.Clear()

	// Receive updates from channel
	for update := range updates {

		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = `/help to see available commands`

		// Log message from user
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			msg.Text = "start"
		case "help":
			msg.Text = `/pic - publish picture in SNE-Life blog`
		case "pic":
			msg.Text = "Send me Photo to publish 🥳"
		}

		if update.Message.Photo != nil {
			HandlePicture(update.Message)
			msg.Text = "Got it. Trying to publish 👌🏻"
		}
		// Answer
		Bot.Send(msg)
	}
}
