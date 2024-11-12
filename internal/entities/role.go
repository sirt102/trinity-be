package entities

import "github.com/google/uuid"

type Role struct {
	PublicID uuid.UUID `gorm:"column:public_id; type:char(36);not null; index:idx_uuid" json:"public_id"`
    RoleName string `gorm:"column:role_name; type:varchar(255);not null" json:"role_name"`
    RoleNote string `gorm:"column:role_note; type:varchar(255);not null" json:"role_note"`
}

func (r *Role) TableName() string {
	return "roles"
}
