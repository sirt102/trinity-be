package entities

const TourTableName = "tours"

type Tour struct {
	Base
	Name           string     `json:"name" binding:"required" gorm:"not null"`
	Price          float64    `json:"price" binding:"required" gorm:"not null"`
	Transportation string     `json:"transportation" binding:"required" gorm:"not null"`
	Status         TourStatus `json:"status" binding:"omitempty,oneof=open cancelled done" gorm:"not null"`
}

type TourStatus string

const (
	OpenTourStatus      TourStatus = "open"
	CancelledTourStatus TourStatus = "cancelled"
	DoneTourStatus      TourStatus = "done"
)
