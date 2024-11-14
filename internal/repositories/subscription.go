package repositories

import (
	"trinity-be/global"
	"trinity-be/internal/entities"

	"github.com/google/uuid"
)

type SubscriptionRepository interface {
	GetSubscriptionByID(id uuid.UUID) (*entities.Subscription, error)
}

type subscriptionRepository struct {
}

func NewSubscriptionRepository() SubscriptionRepository {
	return &subscriptionRepository{}
}

// GetSubscriptionByID implements SubscriptionRepository.
func (c *subscriptionRepository) GetSubscriptionByID(id uuid.UUID) (*entities.Subscription, error) {
	var cp entities.Subscription
	err := global.PostgresQLDB.First(&cp, "subscription_id = (?)", id).Error
	if err != nil {
		return nil, err
	}

	return &cp, nil
}
