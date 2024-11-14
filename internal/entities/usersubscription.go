package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserSubscription struct {
	UserSubscriptionID uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"user_subscription_id"`
	UserID             uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	SubscriptionID     uuid.UUID  `gorm:"type:uuid;not null" json:"subscription_id"`
	SubscribedAt       time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"subscribed_at"`
	ExpiryDate         *time.Time `gorm:"type:timestamp" json:"expiry_date,omitempty"`
}

func (UserSubscription) TableName() string {
	return "user_subscriptions"
}
