package handlers

import (
	"net/http"
	"trinity-be/internal/entities/requests"
	"trinity-be/internal/services"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	subscriptionService services.SubscriptionService
}

func NewSubscriptionHandler(subscriptionService services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

func (vh *SubscriptionHandler) UserRegisterSubscription(c *gin.Context) {
	var req requests.UserRegisterSubscriptionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := vh.subscriptionService.UserRegisterSubscription(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func (vh *SubscriptionHandler) UserConfirmSubscription(c *gin.Context) {
	var req requests.UserPaiedForSubscriptionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = vh.subscriptionService.UserPaiedForSubscription(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
