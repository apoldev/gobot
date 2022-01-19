package telegram

import (
	"github.com/coolphp/gobot/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"os"
)

type Bot struct {
	Bot   *tgbotapi.BotAPI
	Store *store.Store

	handlerMain func(update tgbotapi.Update)
}

func NewBot() *Bot {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT"))

	if err != nil {
		log.Panic(err)
	}

	if os.Getenv("TELEGRAM_BOT_DEBUG") == "true" {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		Bot: bot,
	}

}
