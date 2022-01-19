package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID int64 `json:"user_id"`

	FirstName    string `gorm:"type:varchar(255)" json:"first_name"`
	LastName     string `gorm:"type:varchar(255)" json:"last_name"`
	Username     string `gorm:"type:varchar(255)" json:"username"`
	LanguageCode string `gorm:"type:varchar(255)" json:"username"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
