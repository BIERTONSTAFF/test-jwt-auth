package main

import (
	"os"

	"desq.com.ru/testjwtauth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.RefreshToken{})

	return db, nil
}
