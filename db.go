package main

import (
	c "desq.com.ru/testjwtauth/config"
	"desq.com.ru/testjwtauth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.RefreshToken{})

	return db, nil
}
