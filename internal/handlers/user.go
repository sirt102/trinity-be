package handlers

import (
	"net/http"

	"trinity-be/internal/entities/requests"
	"trinity-be/internal/services"
	"trinity-be/internal/utils/common"

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
	var registerRequest requests.UserRegisterRequest
	err := c.ShouldBindBodyWithJSON(&registerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := uh.userService.Register(&registerRequest)

	c.JSON(200, gin.H{"message": u})
}

func (uh *UserHandler) GetUserByEmail(c *gin.Context) {
	var requestFilter requests.UserFilterRequest
	err := c.ShouldBindJSON(&requestFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := uh.userService.GetUserByEmail(requestFilter.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	u, err := uh.userService.GetUserByID(common.StringToUUID(c.Param("user_id")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}
