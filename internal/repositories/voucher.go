package repositories

import (
	"trinity-be/global"
	"trinity-be/internal/entities"

	"github.com/google/uuid"
)

type VoucherRepository interface {
	CreateNewVoucher(voucher *entities.Voucher) error
	GetVoucherByID(id uuid.UUID) (*entities.Voucher, error)
	GetAll() ([]entities.Voucher, error)
	Delete(id uuid.UUID) error
	GetVoucherByCampaignID(campaignID uuid.UUID) (*entities.Voucher, error)
}

type voucherRepository struct {
}

func NewVoucherRepository() VoucherRepository {
	return &voucherRepository{}
}

func (r *voucherRepository) CreateNewVoucher(voucher *entities.Voucher) error {
	return global.PostgresQLDB.Create(voucher).Error
}

func (r *voucherRepository) GetVoucherByID(voucherID uuid.UUID) (*entities.Voucher, error) {
	var voucher entities.Voucher
	err := global.PostgresQLDB.Where("voucher_id = ?", voucherID).First(&voucher).Error
	if err != nil {
		return nil, err
	}

	return &voucher, nil
}

func (r *voucherRepository) GetAll() ([]entities.Voucher, error) {
	var vouchers []entities.Voucher
	err := global.PostgresQLDB.Find(&vouchers).Error

	return vouchers, err
}

func (r *voucherRepository) Delete(voucherID uuid.UUID) error {
	return global.PostgresQLDB.Where("voucher_id = ?", voucherID).Delete(&entities.Voucher{}).Error
}

func (r *voucherRepository) GetVoucherByCampaignID(campaignID uuid.UUID) (*entities.Voucher, error) {
	var vc entities.Voucher
	err := global.PostgresQLDB.Where("campaign_id = ?", campaignID).First(&vc).Error
	if err != nil {
		return nil, err
	}

	return &vc, nil
}
