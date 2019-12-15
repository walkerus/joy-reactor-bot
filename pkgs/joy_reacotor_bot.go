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
	TelegramBotApi *tgbotapi.BotAPI
	Store Store
}

func (bot *JoyReactorBot) StartUpdatingChatStore()  {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.TelegramBotApi.GetUpdatesChan(u)

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
	lastPostId := ``

	for {
		time.Sleep(30 * time.Second)

		resp, httpGetError := http.Get("http://joyreactor.cc/best")

		if httpGetError != nil {
			log.Println(httpGetError)
			continue
		}

		id, _ := GetPostId(resp.Body)
		resp.Body.Close()

		if lastPostId != id {
			lastPostId = id

			chats, _ := bot.Store.GetChats()
			message := `http://joyreactor.cc/post/` + id

			for _, value := range chats {
				msg := tgbotapi.NewMessage(value, message)

				if _, sendError := bot.TelegramBotApi.Send(msg); sendError != nil {
					log.Println(sendError)
				}
			}
		}
	}
}

func GetPostId(content io.Reader) (string, error)  {
	hrefRe := regexp.MustCompile(`href="/post/\d{1,}"`)
	postIdRe := regexp.MustCompile(`\d{1,}`)

	for true {
		bs := make([]byte, 1014)
		n, err := content.Read(bs)

		substring := hrefRe.Find(bs[:n])
		postId := postIdRe.Find(substring)

		if len(postId) != 0 {
			return string(postId), nil
		}

		if n == 0 || err != nil {
			break
		}
	}

	return ``, errors.New(`post not found`)
}
