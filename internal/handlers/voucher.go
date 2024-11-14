package handlers

import (
	"net/http"
	"trinity-be/internal/entities"
	"trinity-be/internal/services"

	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	voucherService services.VoucherService
}

func NewVoucherHandler(voucherService services.VoucherService) *VoucherHandler {
	return &VoucherHandler{
		voucherService: voucherService,
	}
}

func (vh *VoucherHandler) CreateVoucher(c *gin.Context) {
	var vc entities.Voucher
	err := c.ShouldBindJSON(&vc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = vh.voucherService.CreateNewVoucher(&vc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vc)
}
