package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Medharametlanaveen/jobflow/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	database.AutoMigrate(&models.Job{})

	log.Println("database connected and migrated")
	return database
}
