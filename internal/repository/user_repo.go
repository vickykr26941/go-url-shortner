package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/vickykumar/url_shortner/internal/models"
	"github.com/vickykumar/url_shortner/internal/utils"
	"time"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrInvalidAPIKey      = errors.New("invalid API key")
	ErrDuplicateShortCode = errors.New("short code already exists")
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	if user.APIKey == nil {
		apiKey, err := utils.GenerateAPIKey()
		if err != nil {
			return err
		}
		user.APIKey = &apiKey
	}
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.LastResetDate = now

	query := `INSERT INTO users (
		email, password_hash, api_key, created_at, updated_at,
		last_reset_date, daily_url_count, daily_click_count,
		is_premium, name, last_login_at,id
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.PasswordHash,
		user.APIKey,
		user.CreatedAt,
		user.UpdatedAt,
		user.LastResetDate,
		user.DailyURLCount,
		user.DailyClickCount,
		user.IsPremium,
		user.Name,
		user.LastLoginAt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password_hash,name, api_key, is_premium, created_at, updated_at, last_login_at, daily_url_count, daily_click_count, last_reset_date
				FROM users WHERE id = ?`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.APIKey,
		&user.IsPremium,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLoginAt,
		&user.DailyURLCount,
		&user.DailyClickCount,
		&user.LastResetDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password_hash, name, api_key, is_premium, created_at, updated_at, last_login_at, daily_url_count, daily_click_count, last_reset_date
				FROM users WHERE email = ?`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.APIKey,
		&user.IsPremium,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLoginAt,
		&user.DailyURLCount,
		&user.DailyClickCount,
		&user.LastResetDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetByAPIKey(ctx context.Context, apiKey string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password_hash, name, api_key, is_premium, created_at, updated_at, last_login_at, daily_url_count, daily_click_count, last_reset_date
				FROM users WHERE api_key = ?`

	err := r.db.QueryRowContext(ctx, query, apiKey).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.APIKey,
		&user.IsPremium,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLoginAt,
		&user.DailyURLCount,
		&user.DailyClickCount,
		&user.LastResetDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()

	query := `UPDATE users SET email = ?, password_hash = ?, name = ?, api_key = $4, is_premium = ?, updated_at = ?, last_login_at = ?, daily_url_count = ?, daily_click_count = ?, last_reset_date = ? 
             WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.APIKey,
		user.IsPremium,
		user.UpdatedAt,
		user.LastLoginAt,
		user.DailyURLCount,
		user.DailyClickCount,
		user.LastResetDate,
		user.ID,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			if pqErr.Constraint == "users_email_key" {
				return ErrDuplicateEmail
			}
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (r *userRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	now := time.Now()
	query := `UPDATE users SET last_login_at =?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, now, now, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to update last login: %w", err)
	}
	return nil
}

func (r *userRepository) IncrementDailyCounts(ctx context.Context, userID int64, urls, clicks int) error {

	user, err := r.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user for increment: %w", err)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	lastResetDate := time.Date(user.LastResetDate.Year(), user.LastResetDate.Month(), user.LastResetDate.Day(), 0, 0, 0, 0, user.LastResetDate.Location())

	var newURLCount, newClickCount int
	var newResetDate time.Time

	if today.After(lastResetDate) {
		newURLCount = urls
		newClickCount = clicks
		newResetDate = today
	} else {
		newURLCount = user.DailyURLCount + urls
		newClickCount = user.DailyClickCount + clicks
		newResetDate = user.LastResetDate
	}

	query := `
        UPDATE users SET daily_url_count = ?, daily_click_count = ?, last_reset_date = ?, updated_at = ?
        WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query,
		newURLCount,
		newClickCount,
		newResetDate,
		now,
		userID,
	)

	if err != nil {
		return fmt.Errorf("failed to increment daily counts: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
