package store

import (
	"fmt"
	"github.com/coolphp/gobot/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return nil
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{},
	)

	return db
}
