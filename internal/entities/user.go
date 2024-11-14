package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID           uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"user_id"`
	UserName         string           `gorm:"column:user_name;type:varchar(100);not null" json:"user_name" binding:"required,user_name"`
	Email            string           `gorm:"type:varchar(255);unique;not null" json:"email" binding:"required,email"`
	RegisteredAt     time.Time        `gorm:"autoCreateTime" json:"registered_at"`
	RoleID           uuid.UUID        `gorm:"type:uuid;not null" json:"role_id" binding:"required"`
	Role             Role             `gorm:"foreignKey:RoleID;references:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	UserSubscription UserSubscription `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_subscription"`
	UserVoucher      []UserVoucher    `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_vouchers"`
}

func (u *User) TableName() string {
	return "users"
}
