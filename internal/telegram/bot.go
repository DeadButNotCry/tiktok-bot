package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	groupId int64
}

func NewBot(bot *tgbotapi.BotAPI, groupId int64) *Bot {
	return &Bot{
		bot:     bot,
		groupId: groupId,
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

		if update.Message.Chat.ID == int64(b.groupId) {
			log.Println("Chat message")
			continue
		}

		userIn, err := b.bot.GetChatMember(tgbotapi.ChatConfigWithUser{
			UserID: int(userId),
			ChatID: int64(b.groupId),
		})

		if userIn.Status != "left" && err == nil && update.Message.Video != nil {
			go b.handleVideo(&update)

		} else if userIn.Status != "left" && err == nil {
			go b.handleLink(&update)
		} else if update.Message.IsCommand() {
			go b.handleCommand(&update)
		} else {
			log.Printf("From user %d with no video", userId)
			msg := tgbotapi.NewMessage(userId, "Mistake")
			b.bot.Send(msg)
		}
	}

	return nil
}
