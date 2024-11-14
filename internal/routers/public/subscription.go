package public

import (
	"trinity-be/internal/wire"

	"github.com/gin-gonic/gin"
)

type SubscriptionRouter struct{}

func (tr *SubscriptionRouter) InitSubscriptionRouter(r *gin.RouterGroup) {
	subscriptionHandler := wire.InitSubscriptionRouterHandler()

	// non-auth routes
	subscriptionRouterPublic := r.Group("/subscriptions")
	{
		subscriptionRouterPublic.POST("", subscriptionHandler.UserRegisterSubscription)
		subscriptionRouterPublic.POST("/confirmed", subscriptionHandler.UserConfirmSubscription)
	}

	// auth routes
}
