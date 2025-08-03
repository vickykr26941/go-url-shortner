package models

import (
	"fmt"
	"time"
)

const (
	MAX_DAILY_URL_COUNT   = 1000  // Maximum URLs a user can create in a day
	MAX_DAILY_CLICK_COUNT = 10000 // Maximum clicks a user can have in a day
)

type User struct {
	ID              int64      `json:"id" db:"id"`
	Email           string     `json:"email" db:"email"`
	PasswordHash    string     `json:"-" db:"password_hash"`
	Name            *string    `json:"name" db:"name"`
	APIKey          *string    `json:"api_key" db:"api_key"`
	IsPremium       bool       `json:"is_premium" db:"is_premium"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
	LastLoginAt     *time.Time `json:"last_login_at" db:"last_login_at"`
	DailyURLCount   int        `json:"daily_url_count" db:"daily_url_count"`
	DailyClickCount int        `json:"daily_click_count" db:"daily_click_count"`
	LastResetDate   time.Time  `json:"last_reset_date" db:"last_reset_date"`
}

type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=50"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Name            string `json:"name" validate:"omitempty,min=2,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64        `json:"expires_in"`
}

type UserResponse struct {
	ID          int64      `json:"id"`
	Email       string     `json:"email"`
	Name        *string    `json:"name"`
	IsPremium   bool       `json:"is_premium"`
	CreatedAt   time.Time  `json:"created_at"`
	LastLoginAt *time.Time `json:"last_login_at"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name" validate:"omitempty,min=2,max=100"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"omitempty,min=8,max=50"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type APIKeyResponse struct {
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CanCreateURL() bool {
	if u.DailyURLCount < MAX_DAILY_URL_COUNT {
		return true
	}
	return false
}

func (u *User) GetDailyURLLimit() int {
	return MAX_DAILY_URL_COUNT
}

func (u *User) GetDailyClickLimit() int {
	return MAX_DAILY_CLICK_COUNT
}

func (u *User) ResetDailyCounts() {
	u.DailyURLCount = 0
}

func (u *RegisterRequest) ValidateCreateRequest() error {
	switch {
	case u.Email == "":
		return fmt.Errorf("email is required")
	case u.Name == "":
		return fmt.Errorf("name is required")
	case u.Password == "" || u.ConfirmPassword == "":
		return fmt.Errorf("password and confirm password are required")
	case u.Password != u.ConfirmPassword:
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func (u *LoginRequest) ValidateLoginRequest() error {
	switch {
	case u.Email == "":
		return fmt.Errorf("email is required")
	case u.Password == "":
		return fmt.Errorf("password is required")
	}
	return nil

}
