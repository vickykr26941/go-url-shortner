package repository

import (
	"context"
	"database/sql"
	"github.com/vickykumar/url_shortner/internal/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	// TODO: Insert new user
	// TODO: Handle unique constraint violations (email)
	// TODO: Generate API key if needed
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	// TODO: Retrieve user by ID
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	// TODO: Retrieve user by email
	return nil, nil
}

func (r *userRepository) GetByAPIKey(ctx context.Context, apiKey string) (*models.User, error) {
	// TODO: Retrieve user by API key
	return nil, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	// TODO: Update user information
	// TODO: Update updated_at timestamp
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	// TODO: Delete user
	// TODO: Handle cascade deletion of URLs
	return nil
}

func (r *userRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	// TODO: Update last login timestamp
	return nil
}

func (r *userRepository) IncrementDailyCounts(ctx context.Context, userID int64, urls, clicks int) error {
	// TODO: Increment daily URL and click counts
	// TODO: Reset counts if date has changed
	return nil
}
