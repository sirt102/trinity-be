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

    r.GET("/aa", func(c *gin.Context) {
        global.RedisDB.Set(c, "is_running", "OK", 600)
    })
    v1Group := r.Group("/v1")
    {
        v1Group.GET("/ping")

        publicRouter.InitUserRouter(v1Group)

        // adminRouter.InitAdminRouter(v1Group)
    }


	return r
}