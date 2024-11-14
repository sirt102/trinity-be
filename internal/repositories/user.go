package repositories

import (
	"trinity-be/global"
	"trinity-be/internal/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(userID uuid.UUID) (*entities.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser(user *entities.User) error {
	if err := global.PostgresQLDB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByEmail implements UserRepository.
func (ur *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := global.PostgresQLDB.
		Preload("Role").
		// Preload("Subscription").
		Preload("UserSubscription").
		Preload("UserVoucher").
		First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) GetUserByID(userID uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := global.PostgresQLDB.
		Preload("Role").
		// Preload("Subscription").
		Preload("UserSubscription").
		Preload("UserVoucher").
		First(&user, "user_id =?", userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
