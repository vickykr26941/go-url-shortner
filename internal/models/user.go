package models

import (
	"time"
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
	// TODO: Check if user can create more URLs based on limits
	return false
}

func (u *User) GetDailyURLLimit() int {
	// TODO: Return daily URL creation limit based on user type
	return 0
}

func (u *User) GetDailyClickLimit() int {
	// TODO: Return daily click limit based on user type
	return 0
}

func (u *User) ResetDailyCounts() {
	// TODO: Reset daily counters if date has changed
}
