package requests

import "github.com/google/uuid"

type UserRegisterSubscriptionRequest struct {
	UserID         uuid.UUID `json:"user_id" binding:"required,uuid"`
	SubscriptionID uuid.UUID `json:"subscription_id" binding:"required,uuid"`
}

type UserPaiedForSubscriptionRequest struct {
	UserID         uuid.UUID `json:"user_id" binding:"required,uuid"`
	SubscriptionID uuid.UUID `json:"subscription_id" binding:"required,uuid"`
	UserVoucherID  uuid.UUID `json:"user_voucher_id" binding:"required,uuid"`
	TransactionID  uuid.UUID `json:"transaction_id" binding:"required,uuid"`
}
