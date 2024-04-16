package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/database"
)

func GetUsers(ctx *gin.Context, userRepo database.UserRepository) {
	users, err := userRepo.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
