package requests

type UserFilterRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// UserRegisterRequest - Use temporary structure for registration request
// Replace with Log in with SSO or OAuth when implementing user registration
type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"user_name" binding:"omitempty"`
}
