package db

import (
	"golang_gin_helloworld/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Country{}); err != nil {
		return err
	}
	return nil
}
