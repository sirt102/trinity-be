package handlers

import (
	"trinity-be/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Register(c *gin.Context) {
	u := uh.userService.Register("email", "purpose")

	c.JSON(200, gin.H{"message": u})
}
