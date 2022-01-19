package telegram

import (
	"github.com/coolphp/gobot/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) SaveUserToStore(messageUser *tgbotapi.User) {

	users := repository.NewUserRepository(b.Store)
	users.AddUser(messageUser)

}
