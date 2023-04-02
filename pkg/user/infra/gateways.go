package infra

import (
	"fmt"
	"time"

	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("failed to load .env file: %w", err))
	}
}

func GenerateToken(userID string, expirationTime time.Time) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
