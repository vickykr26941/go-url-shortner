package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/vickykumar/url_shortner/internal/models"
	time2 "time"
)

type urlRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) URLRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) Create(ctx context.Context, url *models.URL) error {
	query := `INSERT INTO urls(id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at, is_active, click_count, is_custom, password_hash
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	time := time2.Now()
	url.CreatedAt = time
	url.UpdatedAt = time

	_, err := r.db.ExecContext(ctx, query,
		url.ID,
		url.ShortCode,
		url.OriginalURL,
		url.UserID,
		url.Title,
		url.Description,
		url.CreatedAt,
		url.UpdatedAt,
		url.ExpiresAt,
		url.IsActive,
		url.ClickCount,
		url.IsCustom,
		url.PasswordHash,
	)

	if err != nil {
		return fmt.Errorf("failed to create url: %w", err)
	}
	return nil
}

func (r *urlRepository) GetByID(ctx context.Context, id int64) (*models.URL, error) {
	url := &models.URL{}
	query := `SELECT id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at, is_active, click_count, is_custom, password_hash FROM urls WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&url.ID,
		&url.ShortCode,
		&url.OriginalURL,
		&url.UserID,
		&url.Title,
		&url.Description,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.ExpiresAt,
		&url.IsActive,
		&url.ClickCount,
		&url.IsCustom,
		&url.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("url with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get url by id: %w", err)
	}

	return url, nil
}

func (r *urlRepository) GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	time := time2.Now()
	url := &models.URL{}
	query := `SELECT id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at, is_active, click_count, is_custom, password_hash FROM urls WHERE short_code = $1 and is_active = true and expires_at > $2`
	err := r.db.QueryRowContext(ctx, query, shortCode, time).Scan(
		&url.ID,
		&url.ShortCode,
		&url.OriginalURL,
		&url.UserID,
		&url.Title,
		&url.Description,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.ExpiresAt,
		&url.IsActive,
		&url.ClickCount,
		&url.IsCustom,
		&url.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("url with id %s not found", shortCode)
		}
		return nil, fmt.Errorf("failed to get url by id: %w", err)
	}
	return url, nil
}

func (r *urlRepository) Update(ctx context.Context, url *models.URL) error {
	// TODO: Update URL in database
	// TODO: Update updated_at timestamp
	return nil
}

func (r *urlRepository) Delete(ctx context.Context, id int64) error {
	// TODO: Delete URL by ID
	// TODO: Handle foreign key constraints
	return nil
}

func (r *urlRepository) List(ctx context.Context, userID int64, req *models.URLListRequest) ([]*models.URL, int64, error) {
	// TODO: List URLs with pagination and filtering
	// TODO: Apply search, tag filters
	// TODO: Apply sorting
	// TODO: Return total count for pagination
	return nil, 0, nil
}

func (r *urlRepository) IncrementClickCount(ctx context.Context, id int64) error {
	// TODO: Atomically increment click count
	return nil
}

func (r *urlRepository) GetExpiredURLs(ctx context.Context) ([]*models.URL, error) {
	// TODO: Get URLs that have expired
	return nil, nil
}

func (r *urlRepository) DeleteExpiredURLs(ctx context.Context) error {
	// TODO: Delete expired URLs
	// TODO: Use batch deletion for performance
	return nil
}
