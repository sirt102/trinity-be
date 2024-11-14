package repositories

import (
	"time"
	"trinity-be/global"
	"trinity-be/internal/entities"

	"github.com/google/uuid"
)

type CampaignRepository interface {
	GetCampaignByID(id uuid.UUID) (*entities.Campaign, error)
	InsertNewCampaign(cp *entities.Campaign) error
	GetRunningCampaigns() ([]entities.Campaign, error)
}

type campaignRepository struct {
}

func NewCampaignRepository() CampaignRepository {
	return &campaignRepository{}
}

// GetCampaignByID implements CampaignRepository.
func (c *campaignRepository) GetCampaignByID(id uuid.UUID) (*entities.Campaign, error) {
	var cp entities.Campaign
	err := global.PostgresQLDB.First(&cp, "campaign_id = (?)", id).Error
	if err != nil {
		return nil, err
	}

	return &cp, nil
}

func (c *campaignRepository) InsertNewCampaign(cp *entities.Campaign) error {
	return global.PostgresQLDB.Create(cp).Error
}

func (c *campaignRepository) GetRunningCampaigns() ([]entities.Campaign, error) {
	var campaigns []entities.Campaign
	now := time.Now()
	err := global.PostgresQLDB.Where("start_date <= ? AND end_date >= ?", now, now).Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}
