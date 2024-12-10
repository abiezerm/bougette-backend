package common

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// for mysql you should read from .env file
	db, err := gorm.Open(sqlite.Open("bougette.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Default().Println("Database connection established")
	return db, nil
}
