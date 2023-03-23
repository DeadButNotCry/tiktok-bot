package main

import (
	"log"
	"os"

	"github.com/deadbutnotcry/tiktok-bot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	os.Mkdir("./videos/", os.ModePerm)
	os.Mkdir("./result/", os.ModePerm)
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	token = "5945815849:AAEz1DnAUCW4G2DGKacwdHvy602T-7l8cZ4"
	if token == "" {
		panic("empty token")
	}
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = false
	bot := telegram.NewBot(botApi)
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
