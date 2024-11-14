package repositories

import (
	"trinity-be/global"
	"trinity-be/internal/entities"
)

type UserSubscriptionRepository interface {
	CreateNewUserSubscription(uv *entities.UserSubscription) error
}

type userSubscriptionRepository struct {
}

func NewUserSubscriptionRepository() UserSubscriptionRepository {
	return &userSubscriptionRepository{}
}

func (r *userSubscriptionRepository) CreateNewUserSubscription(userSubscription *entities.UserSubscription) error {
	return global.PostgresQLDB.Create(userSubscription).Error
}
