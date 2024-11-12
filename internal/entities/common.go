package entities

type Base struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
	Status    string `json:"status" gorm:"not null"`
}

type Status string

const (
	ActiveStatus   Status = "active"
	InactiveStatus Status = "inactive"
)
