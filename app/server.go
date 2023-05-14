package app

import (
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/database"
)

func SetupDatabase() *gorm.DB {
	db := database.ConnectDB()
	database.RunMigrations(db)
	return db
}
