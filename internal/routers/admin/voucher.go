package admin

import (
	"trinity-be/internal/wire"

	"github.com/gin-gonic/gin"
)

type VoucherRouter struct{}

func (tr *VoucherRouter) InitVoucherRouter(r *gin.RouterGroup) {
	campaignHandler := wire.InitVoucherRouterHandler()

	// non-auth routes
	campaignRouterPublic := r.Group("/campaigns")
	{
		campaignRouterPublic.POST("/", campaignHandler.CreateVoucher)
	}
}
