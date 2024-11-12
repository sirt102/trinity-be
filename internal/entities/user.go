package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    PublicID  uuid.UUID `gorm:"column:public_id;type:char(36);not null;index:idx_uuid" json:"public_id"`
    UserName  string    `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`
    FirstName string    `gorm:"column:first_name;type:varchar(255);not null" json:"first_name"`
    LastName  string    `gorm:"column:last_name;type:varchar(255);not null" json:"last_name"`
    Email     string    `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
    IsActive  bool      `gorm:"column:is_active;type:boolean;default:false" json:"is_active"`
    Roles     []Role    `gorm:"many2many:user_roles" json:"roles"`
}

func (u *User) TableName() string {
	return "users"
}
