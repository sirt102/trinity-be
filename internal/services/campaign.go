package services

import (
	"trinity-be/global"
	"trinity-be/internal/entities/requests"
	"trinity-be/internal/repositories"
)

type CampaignService interface {
	CreateNewCampaign(*requests.CreateCampaignRequest) error
}

type campaignService struct {
	campaignRepo repositories.CampaignRepository
	userRepo     repositories.UserRepository
}

func NewCampaignService(campaignRepo repositories.CampaignRepository, userRepo repositories.UserRepository) CampaignService {
	return &campaignService{campaignRepo: campaignRepo, userRepo: userRepo}
}

func (cs *campaignService) CreateNewCampaign(cp *requests.CreateCampaignRequest) error {
	currentUser, err := cs.userRepo.GetUserByID(cp.UserID)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	if !currentUser.Role.AdminPermission {
		return global.NewError("only admin can create new campaigns")
	}

	err = cs.campaignRepo.InsertNewCampaign(&cp.Campaign)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	return nil
}
