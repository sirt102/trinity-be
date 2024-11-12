package initialize

import (
	"fmt"
	"trinity-be/global"
)

func InitializeServer() {
	LoadConfig()
	InitLogger()
	InitPostgresQL()
	InitRedis()
	InitializeRouter().Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
