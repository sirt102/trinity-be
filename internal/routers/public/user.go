package public

import (
	"trinity-be/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (tr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// non di
    userHandler := wire.InitUserRouterHandler()
  
	
	// non-auth routes
	userRouterPublic := r.Group("/users")
	{
		userRouterPublic.GET("/" )
		userRouterPublic.GET("/:id")
        userRouterPublic.POST("/register", userHandler.Register)
	}

	// auth routes
}
