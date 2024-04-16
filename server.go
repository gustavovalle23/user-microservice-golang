package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	generalinfra "github.com/gustavovalle23/user-microservice-golang/pkg/general_infra"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/database"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/infra"
	_ "github.com/lib/pq"
)

func main() {
	dbInfo := generalinfra.GetConfig()

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	r := gin.Default()

	userRepo := database.UserRepository{DB: db}

	r.GET("/users", func(ctx *gin.Context) {
		infra.GetUsers(ctx, userRepo)
	})

	r.GET("/users/{int}", func(ctx *gin.Context) {
		infra.GetUsers(ctx, userRepo)
	})

	r.Run("localhost:8000")
}
