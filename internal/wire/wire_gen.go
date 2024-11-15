// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"trinity-be/internal/handlers"
	"trinity-be/internal/repositories"
	"trinity-be/internal/services"
)

// Injectors from campaign.go:

func InitCampaignRouterHandler() *handlers.CampaignHandler {
	campaignRepository := repositories.NewCampaignRepository()
	userRepository := repositories.NewUserRepository()
	campaignService := services.NewCampaignService(campaignRepository, userRepository)
	campaignHandler := handlers.NewCampaignHandler(campaignService)
	return campaignHandler
}

// Injectors from subscription.go:

func InitSubscriptionRouterHandler() *handlers.SubscriptionHandler {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	campaignRepository := repositories.NewCampaignRepository()
	voucherRepository := repositories.NewVoucherRepository()
	userVoucherRepository := repositories.NewUserVoucherRepository()
	transactionRepository := repositories.NewTransactionRepository()
	userSubscriptionRepository := repositories.NewUserSubscriptionRepository()
	subscriptionService := services.NewSubscriptionService(subscriptionRepository, campaignRepository, voucherRepository, userVoucherRepository, transactionRepository, userSubscriptionRepository)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)
	return subscriptionHandler
}

// Injectors from user.go:

func InitUserRouterHandler() *handlers.UserHandler {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}

// Injectors from voucher.go:

func InitVoucherRouterHandler() *handlers.VoucherHandler {
	voucherRepository := repositories.NewVoucherRepository()
	voucherService := services.NewVoucherService(voucherRepository)
	voucherHandler := handlers.NewVoucherHandler(voucherService)
	return voucherHandler
}
