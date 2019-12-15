package pkgs

import (
	"errors"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

type JoyReactorBot struct {
	TelegramBotAPI *tgbotapi.BotAPI
	Store          Store
}

func (bot *JoyReactorBot) StartUpdatingChatStore() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.TelegramBotAPI.GetUpdatesChan(u)

	if err != nil {
		log.Println(err)
		return
	}

	for update := range updates {
		if update.Message.Text != `/start` { // ignore any non-Message Updates
			continue
		}

		if addChatError := bot.Store.AddChat(update.Message.Chat.ID); addChatError != nil {
			log.Println(addChatError)
		}
	}
}

func (bot *JoyReactorBot) StartMailing() {
	lastPostID := ``

	for {
		time.Sleep(30 * time.Second)

		resp, httpGetError := http.Get("http://joyreactor.cc/best")

		if httpGetError != nil {
			log.Println(httpGetError)
			continue
		}

		id, _ := GetPostID(resp.Body)
		resp.Body.Close()

		if lastPostID != id {
			lastPostID = id

			chats, _ := bot.Store.GetChats()
			message := `http://joyreactor.cc/post/` + id

			for _, value := range chats {
				msg := tgbotapi.NewMessage(value, message)

				if _, sendError := bot.TelegramBotAPI.Send(msg); sendError != nil {
					log.Println(sendError)
				}
			}
		}
	}
}

func GetPostID(content io.Reader) (string, error) {
	hrefRe := regexp.MustCompile(`href="/post/\d{1,}"`)
	postIDRe := regexp.MustCompile(`\d{1,}`)

	for true {
		bs := make([]byte, 1014)
		n, err := content.Read(bs)

		substring := hrefRe.Find(bs[:n])
		postID := postIDRe.Find(substring)

		if len(postID) != 0 {
			return string(postID), nil
		}

		if n == 0 || err != nil {
			break
		}
	}

	return ``, errors.New(`post not found`)
}
