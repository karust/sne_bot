package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	//pwd, _ := os.Getwd()
	fmt.Println(os.Getenv("TELEGRAM_APITOKEN"))

	token, err := ioutil.ReadFile("./token.txt")
	if err != nil {
		print("Can't read token file: ", err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		print("Can't load API token", err)
		return
	}

	bot.Debug = true // Has the library display every request and response.
}
