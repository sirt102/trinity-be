package services

import (
	"trinity-be/global"
	"trinity-be/internal/entities"
	"trinity-be/internal/repositories"
)

type VoucherService interface {
	CreateNewVoucher(*entities.Voucher) error
}

type voucherService struct {
	voucherRepo repositories.VoucherRepository
}

func NewVoucherService(voucherRepo repositories.VoucherRepository) VoucherService {
	return &voucherService{voucherRepo: voucherRepo}
}

// CreateNewVoucer implements VoucherService.
func (v *voucherService) CreateNewVoucher(vc *entities.Voucher) error {
	err := v.voucherRepo.CreateNewVoucher(vc)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	return nil
}
