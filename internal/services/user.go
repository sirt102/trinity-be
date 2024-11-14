package services

import (
	"log"
	"strings"
	"trinity-be/global"
	"trinity-be/internal/entities"
	"trinity-be/internal/entities/requests"
	"trinity-be/internal/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	Register(*requests.UserRegisterRequest) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(userID uuid.UUID) (*entities.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) Register(req *requests.UserRegisterRequest) error {
	targetUser, err := us.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	if targetUser != nil {
		return global.NewError("Email already registered")
	}

	if req.UserName == "" {
		req.UserName = strings.Split(req.Email, "@")[0]
	}
	var newUser = &entities.User{
		Email:    req.Email,
		UserName: req.UserName,
	}

	err = us.userRepo.CreateUser(newUser)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	return nil
}

func (us *userService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		global.LogError(err, "")
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserByID(userID uuid.UUID) (*entities.User, error) {
	log.Println("User ID", userID)
	user, err := us.userRepo.GetUserByID(userID)
	if err != nil {
		global.LogError(err, "")
		return nil, err
	}

	return user, nil
}
