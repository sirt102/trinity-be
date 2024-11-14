package entities

import (
	"time"

	"github.com/google/uuid"
)

// VoucherUser struct to represent the VoucherUser table
type UserVoucher struct {
	UserVoucherID uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"user_voucher_id"`
	VoucherID     uuid.UUID  `gorm:"type:uuid;not null" json:"voucher_id"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	RedeemedAt    *time.Time `gorm:"type:timestamp" json:"redeemed_at,omitempty"`
	Voucher       Voucher    `gorm:"foreignKey:VoucherID;references:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"voucher"`
}

func (UserVoucher) TableName() string {
	return "user_vouchers"
}
