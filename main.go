package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"joy-reactor/pkgs"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile(os.Getenv(`BOT_LOG_FILE`), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	tgBotApi, err := tgbotapi.NewBotAPI(os.Getenv(`BOT_TOKEN`))

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Authorized on account %s", tgBotApi.Self.UserName)

	s := pkgs.Store{FileName: os.Getenv(`BOT_STORE_FILE`)}
	joyReactorBot := pkgs.JoyReactorBot{TelegramBotApi: tgBotApi, Store: s}

	go joyReactorBot.StartUpdatingChatStore()
	joyReactorBot.StartMailing()
}
