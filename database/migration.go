package database

import (
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})
}
