package repositories

import (
	"time"
	"trinity-be/global"
	"trinity-be/internal/entities"

	"github.com/google/uuid"
)

type UserVoucherRepository interface {
	CreateNewUserVoucher(uv *entities.UserVoucher) error
	UsedUserVoucher(id uuid.UUID) error
	GetUserVoucherByID(id uuid.UUID) (*entities.UserVoucher, error)
}

type userVoucherRepository struct {
}

func NewUserVoucherRepository() UserVoucherRepository {
	return &userVoucherRepository{}
}

func (r *userVoucherRepository) CreateNewUserVoucher(userVoucher *entities.UserVoucher) error {
	return global.PostgresQLDB.Create(userVoucher).Error
}

func (r *userVoucherRepository) UsedUserVoucher(id uuid.UUID) error {
	return global.PostgresQLDB.Model(&entities.UserVoucher{}).Where("user_voucher_id = ?", id).Update("redeemed_at", time.Now()).Error
}

func (r *userVoucherRepository) GetUserVoucherByID(id uuid.UUID) (*entities.UserVoucher, error) {
	var uv entities.UserVoucher
	err := global.PostgresQLDB.First(&uv, "user_voucher_id =?", id).Error
	if err != nil {
		return nil, err
	}

	return &uv, nil
}
