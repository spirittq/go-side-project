package db

import (
	"sideq/internal/entity"
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Gorm struct {
	DB *gorm.DB
}

func New(url string) (*Gorm, error) {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.Example{})
	return &Gorm{
		DB: db,
	}, nil
}
