package admin

import (
	"trinity-be/internal/wire"

	"github.com/gin-gonic/gin"
)

type CampaignRouter struct{}

func (tr *CampaignRouter) InitCampaignRouter(r *gin.RouterGroup) {
	campaignHandler := wire.InitCampaignRouterHandler()

	// non-auth routes
	campaignRouterPublic := r.Group("/campaigns")
	{
		campaignRouterPublic.POST("/", campaignHandler.CreateCampaign)
	}

	// auth routes
	// Will implement once authentication is implemented - in the future
}
