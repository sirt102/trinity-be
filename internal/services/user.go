package services

import (
	"trinity-be/internal/repositories"
)

type UserService interface {
    Register(email, purpose string) int
}

type userService struct {
    userRepo *repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo: &userRepo}
}

func (us *userService) Register(email, purpose string) int {
    return 1
}