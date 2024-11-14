package entities

import "github.com/google/uuid"

type Subscription struct {
	SubscriptionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"subscription_id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Price          float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Level          int       `gorm:"type:int;not null" json:"level"`
	Description    string    `gorm:"type:text" json:"description"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
