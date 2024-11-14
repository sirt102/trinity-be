package entities

import (
	"time"

	"github.com/google/uuid"
)

type StatusPayment string

const (
	PaymentStatusPending  StatusPayment = "pending"
	PaymentStatusSuccess  StatusPayment = "success"
	PaymentStatusFailed   StatusPayment = "failed"
	PaymentStatusCanceled StatusPayment = "canceled"
)

type Transaction struct {
	TransactionID   uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"transaction_id"`
	UserID          uuid.UUID     `gorm:"type:uuid;not null" json:"user_id" binding:"required"`
	VoucherID       uuid.UUID     `gorm:"type:uuid" json:"voucher_id"`
	TransactionDate time.Time     `gorm:"autoCreateTime" json:"transaction_date"`
	Amount          float64       `gorm:"type:decimal(10,2);not null" json:"amount" binding:"required,gt=0"`
	FinalAmount     float64       `gorm:"type:decimal(10,2);not null" json:"final_amount" binding:"required,gt=0"`
	SubscriptionID  uuid.UUID     `gorm:"type:uuid;not null" json:"subscription_id"`
	StatusPayment   StatusPayment `gorm:"type:varchar(20);not null" json:"status_payment"`

	User         User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Voucher      Voucher      `gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Subscription Subscription `gorm:"foreignKey:SubscriptionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (*Transaction) TableName() string {
	return "transactions"
}
