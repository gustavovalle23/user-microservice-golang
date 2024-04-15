package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	IsStaff  bool   `json:"is_staff"`
}

func main() {
	r := gin.Default()

	users := []User{
		{ID: 1, Name: "Gustavo Valle", IsActive: true, IsStaff: true},
	}

	r.GET("/users/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, users)
	})

	r.Run("localhost:8000")
}
