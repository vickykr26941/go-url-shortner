package repository

import (
	"context"
	models2 "github.com/vickykumar/url_shortner/internal/models"
	"time"
)

type URLRepository interface {
	Create(ctx context.Context, url *models2.URL) error
	GetByID(ctx context.Context, id int64) (*models2.URL, error)
	GetByShortCode(ctx context.Context, shortCode string) (*models2.URL, error)
	Update(ctx context.Context, url *models2.URL) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, userID int64, req *models2.URLListRequest) ([]*models2.URL, int64, error)
	IncrementClickCount(ctx context.Context, id int64) error
	GetExpiredURLs(ctx context.Context) ([]*models2.URL, error)
	DeleteExpiredURLs(ctx context.Context) error
}

type UserRepository interface {
	Create(ctx context.Context, user *models2.User) error
	GetByID(ctx context.Context, id int64) (*models2.User, error)
	GetByEmail(ctx context.Context, email string) (*models2.User, error)
	GetByAPIKey(ctx context.Context, apiKey string) (*models2.User, error)
	Update(ctx context.Context, user *models2.User) error
	Delete(ctx context.Context, id int64) error
	UpdateLastLogin(ctx context.Context, id int64) error
	IncrementDailyCounts(ctx context.Context, userID int64, urls, clicks int) error
}

type AnalyticsRepository interface {
	Create(ctx context.Context, analytics *models2.Analytics) error
	GetByURLID(ctx context.Context, urlID int64, dateRange *models2.AnalyticsDateRange) ([]*models2.Analytics, error)
	GetSummary(ctx context.Context, urlID int64, dateRange *models2.AnalyticsDateRange) (*models2.AnalyticsSummary, error)
	GetUserSummary(ctx context.Context, userID int64, dateRange *models2.AnalyticsDateRange) (*models2.AnalyticsSummary, error)
	DeleteOldRecords(ctx context.Context, olderThan time.Time) error
}

type TagRepository interface {
	CreateTags(ctx context.Context, urlID int64, tags []string) error
	GetTagsByURLID(ctx context.Context, urlID int64) ([]string, error)
	DeleteTagsByURLID(ctx context.Context, urlID int64) error
	GetPopularTags(ctx context.Context, userID int64, limit int) ([]string, error)
}
