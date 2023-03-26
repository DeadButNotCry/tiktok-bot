package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		userId := update.Message.Chat.ID
		userIn, err := b.bot.GetChatMember(tgbotapi.ChatConfigWithUser{
			UserID: int(userId),
			ChatID: -1001923600726,
		})

		if update.Message.Chat.ID == -1001923600726 {
			log.Println("Chat message")
			continue
		}

		if userIn.Status != "left" && err == nil && update.Message.Video != nil {
			go b.handleVideo(&update)
		} else {
			log.Println(fmt.Sprintf("From user %s"))
			msg := tgbotapi.NewMessage(userId, "Mistake")
			b.bot.Send(msg)
		}
	}

	return nil
}
