package initialize

import (
	"github.com/gin-gonic/gin"
)

func InitializeServer() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitPostgresQL()
	InitRedis()
	return InitializeRouter()
}
