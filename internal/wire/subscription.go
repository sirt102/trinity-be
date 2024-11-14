//go:build wireinject

package wire

import (
	"trinity-be/internal/handlers"
	"trinity-be/internal/repositories"
	"trinity-be/internal/services"

	"github.com/google/wire"
)

func InitSubscriptionRouterHandler() *handlers.SubscriptionHandler {
	wire.Build(
		repositories.NewSubscriptionRepository,
		repositories.NewCampaignRepository,
		repositories.NewVoucherRepository,
		repositories.NewUserVoucherRepository,
		repositories.NewTransactionRepository,
		repositories.NewUserSubscriptionRepository,
		services.NewSubscriptionService,
		handlers.NewSubscriptionHandler,
	)

	return new(handlers.SubscriptionHandler)
}
