package store

import (
	"gorm.io/gorm"
)

type Store struct {
	Db *gorm.DB
}

func NewStore() *Store {

	db := Connect()

	s := &Store{
		Db: db,
	}

	return s
}
