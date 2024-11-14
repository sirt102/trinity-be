package requests

import (
	"trinity-be/internal/entities"

	"github.com/google/uuid"
)

type CreateCampaignRequest struct {
	UserID   uuid.UUID         `json:"user_id" binding:"required"`
	Campaign entities.Campaign `json:"campaign,inline"`
}
