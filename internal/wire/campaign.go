//go:build wireinject

package wire

import (
	"trinity-be/internal/handlers"
	"trinity-be/internal/repositories"
	"trinity-be/internal/services"

	"github.com/google/wire"
)

func InitCampaignRouterHandler() *handlers.CampaignHandler {
	wire.Build(
		repositories.NewCampaignRepository,
		repositories.NewUserRepository,
		services.NewCampaignService,
		handlers.NewCampaignHandler,
	)

	return new(handlers.CampaignHandler)
}
