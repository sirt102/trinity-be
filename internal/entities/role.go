package entities

import "github.com/google/uuid"

type Role struct {
	RoleID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"role_id"`
	Name            string    `gorm:"type:varchar(50);unique;not null" json:"name" binding:"omitempty,oneof=admin user"`
	AdminPermission bool      `gorm:"type:boolean;default:false" json:"admin_permission" binding:"-"`
}

func (r *Role) TableName() string {
	return "roles"
}
