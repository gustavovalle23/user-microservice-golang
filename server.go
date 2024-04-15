package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dbInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbHost, dbName, dbSSLMode)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	r := gin.Default()

	userRepo := database.UserRepository{DB: db}

	r.GET("/users", func(ctx *gin.Context) {
		users, err := userRepo.GetUsers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	r.Run("localhost:8000")
}
