package entities

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	CampaignID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"campaign_id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name" binding:"required,min=3"`
	StartDate      time.Time `gorm:"type:date;not null" json:"start_date" binding:"required"`
	EndDate        time.Time `gorm:"type:date;not null" json:"end_date" binding:"required,gtfield=StartDate"`
	MaxRedemptions int       `gorm:"type:int;not null" json:"max_redemptions" binding:"required,gte=1"`
	Available      int       `gorm:"type:int;not null" json:"available" binding:"required,gte=0"`
}

func (*Campaign) TableName() string {
	return "campaigns"
}
