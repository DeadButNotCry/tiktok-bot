package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleVideo(update *tgbotapi.Update) {
	videoId := update.Message.Video.FileID
	log.Println(videoId)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, videoId)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
		return
	}

}
