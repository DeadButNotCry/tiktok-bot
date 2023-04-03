package main

import (
	"log"
	"os"
	"strconv"

	"github.com/deadbutnotcry/tiktok-bot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	os.Mkdir("./videos/", os.ModePerm)
	os.Mkdir("./result/", os.ModePerm)
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	groupId, err := strconv.ParseInt(os.Getenv("GROUP_ID"), 10, 64)
	if token == "" || err != nil {
		panic(err)
	}
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = false
	bot := telegram.NewBot(botApi, groupId)
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
