package initialize

import (
	"trinity-be/global"
	"trinity-be/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// adminRouter := routers.RouterGroupApp.Admin
	publicRouter := routers.RouterGroupApp.Public
	v1Group := r.Group("/v1")
	{
		publicRouter.InitUserRouter(v1Group)
	}

	adminRouter := routers.RouterGroupApp.Admin
	v1GroupAdmin := r.Group("/admin/v1")
	{
		adminRouter.InitCampaignRouter(v1GroupAdmin)
	}

	return r
}
