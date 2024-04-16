package generalinfra

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbHost, dbName, dbSSLMode)

}
