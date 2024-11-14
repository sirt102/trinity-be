package public

import (
	"trinity-be/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (tr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userHandler := wire.InitUserRouterHandler()

	// non-auth routes
	userRouterPublic := r.Group("/users")
	{
		userRouterPublic.GET("/details", userHandler.GetUserByEmail)
		userRouterPublic.GET("/details/:user_id", userHandler.GetUserByID)
		userRouterPublic.POST("/register", userHandler.Register)
	}

	// auth routes
}
