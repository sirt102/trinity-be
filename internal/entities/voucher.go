package entities

import (
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	VoucherID    uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"voucher_id"`
	CampaignID   uuid.UUID  `gorm:"type:uuid;not null" json:"campaign_id" binding:"required"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null" json:"user_id" binding:"required"`
	IsValid      bool       `gorm:"default:true" json:"is_valid"`
	ExpiryDate   *time.Time `gorm:"type:date" json:"expiry_date,omitempty"`
	RedeemedAt   *time.Time `gorm:"type:timestamp" json:"redeemed_at,omitempty"`
	DiscountRate float64    `gorm:"type:decimal(5,2)" json:"discount_rate"`

	Campaign Campaign `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User     User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (*Voucher) TableName() string {
	return "vouchers"
}
