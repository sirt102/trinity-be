package handlers

import (
	"net/http"
	"trinity-be/internal/entities/requests"
	"trinity-be/internal/services"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService services.CampaignService
}

func NewCampaignHandler(campaignService services.CampaignService) *CampaignHandler {
	return &CampaignHandler{
		campaignService: campaignService,
	}
}

func (vh *CampaignHandler) CreateCampaign(c *gin.Context) {
	var cp requests.CreateCampaignRequest
	err := c.ShouldBindJSON(&cp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = vh.campaignService.CreateNewCampaign(&cp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
