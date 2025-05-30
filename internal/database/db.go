package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Public variable (UpperCase) that holds the database connection
var DB *gorm.DB

func ConnectDB() {
	var err error

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		panic("failed to load .env file: " + err.Error())
	}

	// Construct the Data Source Name (DSN) for PostgreSQL connection
	// SprintF formats a string with placeholders %s
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	//gorm.Open trys to connect to the database and return two values:
	// 1. A pointer to the gorm.DB instance
	// 2. An error if the connection fails
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
}
