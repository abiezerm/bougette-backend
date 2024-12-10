package main

import (
	"bouguette/common"
	"bouguette/internal/models"
	"log"
)

func main() {
	db, err := common.NewSqliteConnection()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic(err)
	}

	log.Println("Migration completed")
}
