package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=vincentdizon dbname=mydb sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
