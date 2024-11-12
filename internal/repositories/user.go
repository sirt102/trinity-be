package repositories

type UserRepository interface {
	GetUser() string
	GetUserByEmail(email string) string
}

type userRepository struct {
	// db *gorm.DB
}

// GetUserByEmail implements UserRepository.
func (ur *userRepository) GetUserByEmail(email string) string {
	panic("unimplemented")
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUser() string {
	return "User"
}
