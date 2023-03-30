package main

import (
	"fmt"
	"github.com/deadbutnotcry/tiktok-bot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"runtime"
)

func init() {
	os.Mkdir("./videos/", os.ModePerm)
	os.Mkdir("./result/", os.ModePerm)
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	if runtime.GOOS == "windows" {
		log.Print("Pls input bot token")
		fmt.Scanln(&token)
	}
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
