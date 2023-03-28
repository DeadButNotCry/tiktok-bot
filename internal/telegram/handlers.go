package telegram

import (
	"github.com/deadbutnotcry/tiktok-bot/internal/filetransfer"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/deadbutnotcry/tiktok-bot/internal/uniqueizer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Bot) handleVideo(update *tgbotapi.Update) {
	videoId := update.Message.Video.FileID
	chatId := update.Message.Chat.ID
	log.Println(videoId, update.Message.Video.MimeType)
	videoType := strings.Split(update.Message.Video.MimeType, "/")[1]
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, videoId)
	_, err := b.bot.Send(msg)
	check(err)
	fileConfig := tgbotapi.FileConfig{FileID: videoId}
	file, err := b.bot.GetFile(fileConfig)
	if err != nil {
		msg := tgbotapi.NewMessage(chatId, "File is too big.")
		b.bot.Send(msg)
	}

	localFileName := file.FileID + "." + videoType
	path := filepath.Join("./videos/", localFileName)
	out, err := os.Create(path)
	check(err)
	defer out.Close()
	link := file.Link(b.bot.Token)
	log.Println(link)

	resp, err := http.Get(link)
	check(err)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	check(err)
	err = uniqueizer.DoUnique(path, "./result/"+localFileName)
	check(err)

	bf, err := os.ReadFile("./result/" + localFileName)
	check(err)
	videoFile := tgbotapi.FileBytes{Name: localFileName, Bytes: bf}
	video := tgbotapi.NewVideoUpload(chatId, videoFile)
	_, err = b.bot.Send(video)
	check(err)

}

func (b *Bot) handleLink(update *tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	link := update.Message.Text
	errMsg := tgbotapi.NewMessage(chatId, "Try again or write to coder @deadbutnotcry")
	if link == "" || !strings.HasSuffix(link, "#link") || !strings.HasPrefix(link, "https://filetransfer.io/data-package/") {
		b.bot.Send(errMsg)
		return
	}
	localFileName := filetransfer.FileTransferIoDownload(link)
	uniqueizer.DoUnique("./videos/"+localFileName, "./result/"+localFileName)
	//bf, err := os.ReadFile("./result/" + localFileName)
	//check(err)
	//videoFile := tgbotapi.FileBytes{Name: localFileName, Bytes: bf}
	//video := tgbotapi.NewVideoUpload(chatId, videoFile)
	//_, err = b.bot.Send(video)
	//check(err)
	link = filetransfer.UploadToAnonfiles("./result/" + localFileName)
	msg := tgbotapi.NewMessage(chatId, link)
	_, err := b.bot.Send(msg)
	check(err)
}

func (b *Bot) handleCommand(update *tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatId, "По развитию,вопросам и идеям писать @deadbutnotcry")
	b.bot.Send(msg)
}
