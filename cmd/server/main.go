package main

import (
	"fmt"
	"trinity-be/global"
	"trinity-be/internal/initialize"

	_ "trinity-be/cmd/swag/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Trinity Server
// @version         1.0.0
// @description     This is a sample Trinity server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   TriNM
// @contact.email  trisnm102@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /v1
// @schema http
// @securityDefinitions.basic  BasicAuth
func main() {
	r := initialize.InitializeServer()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
