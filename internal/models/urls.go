package models

import (
	"time"
)

type URL struct {
	ID           int64      `json:"id" db:"id"`
	ShortCode    string     `json:"short_code" db:"short_code"`
	OriginalURL  string     `json:"original_url" db:"original_url"`
	UserID       *int64     `json:"user_id" db:"user_id"`
	Title        *string    `json:"title" db:"title"`
	Description  *string    `json:"description" db:"description"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	ExpiresAt    *time.Time `json:"expires_at" db:"expires_at"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	IsCustom     bool       `json:"is_custom" db:"is_custom"`
	PasswordHash *string    `json:"-" db:"password_hash"`
	ClickCount   int64      `json:"click_count" db:"click_count"`
}

type CreateURLRequest struct {
	OriginalURL string     `json:"original_url" validate:"required,url"`
	CustomCode  *string    `json:"custom_code" validate:"omitempty,alphanum,min=3,max=10"`
	Title       *string    `json:"title" validate:"omitempty,max=255"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	ExpiresAt   *time.Time `json:"expires_at"`
	Password    *string    `json:"password" validate:"omitempty,min=4,max=50"`
	Tags        []string   `json:"tags" validate:"omitempty,dive,alphanum,max=50"`
}

type UpdateURLRequest struct {
	Title       *string    `json:"title" validate:"omitempty,max=255"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	ExpiresAt   *time.Time `json:"expires_at"`
	IsActive    *bool      `json:"is_active"`
	Tags        []string   `json:"tags" validate:"omitempty,dive,alphanum,max=50"`
}

type URLResponse struct {
	ID          int64      `json:"id"`
	ShortCode   string     `json:"short_code"`
	ShortURL    string     `json:"short_url"`
	OriginalURL string     `json:"original_url"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
	ClickCount  int64      `json:"click_count"`
	IsCustom    bool       `json:"is_custom"`
	Tags        []string   `json:"tags"`
}

type URLListRequest struct {
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
	Search   string `json:"search" query:"search"`
	Tag      string `json:"tag" query:"tag"`
	SortBy   string `json:"sort_by" query:"sort_by"`
	SortDir  string `json:"sort_dir" query:"sort_dir"`
}

func (u *URL) IsExpired() bool {
	// TODO: Check if URL is expired
	return false
}

func (u *URL) HasPassword() bool {
	// TODO: Check if URL has password protection
	return false
}

func (u *URL) ValidatePassword(password string) bool {
	// TODO: Validate password against hash
	return false
}
