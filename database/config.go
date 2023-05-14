package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/models"
)

func SetupDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}

func RunMigrations() {
	db, err := SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	defer db.Close()

	// Menjalankan migrasi untuk model User
	db.AutoMigrate(&models.User{})

	// Menjalankan migrasi untuk model Photo
	db.AutoMigrate(&models.Photo{})

	log.Println("Database migration completed")
}
