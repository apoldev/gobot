package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func (b *Bot) RegisterHandler(f func(update tgbotapi.Update)) {
	b.handlerMain = f
}

func (b *Bot) StartGetUpdates() {

	if b.handlerMain == nil {
		log.Warn("handlerMain not set")
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {
			// Если указан юзер - сохраним юзера
			if update.Message.From != nil {
				b.SaveUserToStore(update.Message.From)
			}
		}

		b.handlerMain(update)

	}

}
