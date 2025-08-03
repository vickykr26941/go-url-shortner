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
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

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
	query := `SELECT id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at, is_active, click_count, is_custom, password_hash FROM urls WHERE id = ?`
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
	query := `SELECT id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at, is_active, click_count, is_custom, password_hash FROM urls WHERE short_code = ? AND is_active = true AND expires_at > ?`
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
			return nil, fmt.Errorf("url with short_code %s not found", shortCode)
		}
		return nil, fmt.Errorf("failed to get url by short_code: %w", err)
	}

	return url, nil
}
func (r *urlRepository) Update(ctx context.Context, url *models.URL) error {
	query := `UPDATE urls 
			  SET short_code = ?, original_url = ?, title = ?, description = ?, updated_at = ?, 
				  expires_at = ?, is_active = ?, click_count = ?, is_custom = ?, password_hash = ?
			  WHERE id = ?`

	url.UpdatedAt = time2.Now()

	_, err := r.db.ExecContext(ctx, query,
		url.ShortCode,
		url.OriginalURL,
		url.Title,
		url.Description,
		url.UpdatedAt,
		url.ExpiresAt,
		url.IsActive,
		url.ClickCount,
		url.IsCustom,
		url.PasswordHash,
		url.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update url: %w", err)
	}
	return nil
}

func (r *urlRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM urls WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete url: %w", err)
	}
	return nil
}

func (r *urlRepository) List(ctx context.Context, userID int64, req *models.URLListRequest) ([]*models.URL, int64, error) {
	var urls []*models.URL

	query := `SELECT id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at, 
			  is_active, click_count, is_custom, password_hash 
			  FROM urls WHERE user_id = ? LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, userID, req.PageSize, req.Page)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list urls: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		url := &models.URL{}
		err := rows.Scan(
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
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		urls = append(urls, url)
	}

	countQuery := `SELECT COUNT(*) FROM urls WHERE user_id = ?`
	var count int64
	err = r.db.QueryRowContext(ctx, countQuery, userID).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count urls: %w", err)
	}

	return urls, count, nil
}

func (r *urlRepository) IncrementClickCount(ctx context.Context, id int64) error {
	query := `UPDATE urls SET click_count = click_count + 1 WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to increment click count: %w", err)
	}
	return nil
}

func (r *urlRepository) GetExpiredURLs(ctx context.Context) ([]*models.URL, error) {
	query := `SELECT id, short_code, original_url, user_id, title, description, created_at, updated_at, expires_at,
			  is_active, click_count, is_custom, password_hash 
			  FROM urls WHERE expires_at <= ? AND is_active = true`

	now := time2.Now()
	rows, err := r.db.QueryContext(ctx, query, now)
	if err != nil {
		return nil, fmt.Errorf("failed to get expired urls: %w", err)
	}
	defer rows.Close()

	var urls []*models.URL
	for rows.Next() {
		url := &models.URL{}
		err := rows.Scan(
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
			return nil, fmt.Errorf("failed to scan expired url: %w", err)
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (r *urlRepository) DeleteExpiredURLs(ctx context.Context) error {
	query := `DELETE FROM urls WHERE expires_at <= ? AND is_active = true`
	now := time2.Now()

	_, err := r.db.ExecContext(ctx, query, now)
	if err != nil {
		return fmt.Errorf("failed to delete expired urls: %w", err)
	}
	return nil
}
