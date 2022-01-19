package main

import (
	"github.com/coolphp/gobot/server"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func main() {

	s := server.New()
	s.Configure()

	s.Bot.RegisterHandler(func(update tgbotapi.Update) {

		if update.Message != nil {

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := s.Bot.Bot.Send(msg); err != nil {
				log.Warn(err)
			}

		}

	})

	s.Start()

}
