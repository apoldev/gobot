package repository

import (
	"github.com/apoldev/gobot/models"
	"github.com/apoldev/gobot/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserRepository struct {
	store *store.Store
}

func NewUserRepository(store2 *store.Store) *UserRepository {
	return &UserRepository{
		store: store2,
	}
}

func (u *UserRepository) AddUser(userMessage *tgbotapi.User) *models.User {

	var user models.User

	u.store.Db.Model(&models.User{}).Where("user_id = ?", userMessage.ID).First(&user)

	// Create if not exist
	if user.ID <= 0 {
		user.UserID = userMessage.ID
		user.Username = userMessage.UserName
		user.FirstName = userMessage.FirstName
		user.LastName = userMessage.LastName
		user.LanguageCode = userMessage.LanguageCode

		u.store.Db.Create(&user)

	}

	return &user

}
