package models

// TokenResponse represents the response with access token and user ID
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       uint   `json:"user_id"`
	RoleID       uint   `json:"role_id"`
}

// RefreshTokenResponse represents the response with access token and user ID
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// ErrorResponse represents an error message response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DefaultResponse represents a default message response
type DefaultResponse struct {
	Message string `json:"message"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`

	Salary        float32 `gorm:"not null"`
	Position      string  `json:"position"`
	Plan          uint    `json:"plan" gorm:"default:0"`
	SalaryProject uint    `json:"salary_project" gorm:"default:0"`
	PlaceWork     string  `json:"place_work"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
