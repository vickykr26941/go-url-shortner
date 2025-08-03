package service

import (
	"context"
	"github.com/vickykumar/url_shortner/internal/database"
	"strconv"
	"time"
)

type CacheService interface {
	// URL Caching

	SetURL(ctx context.Context, shortCode, originalURL string, ttl time.Duration) error
	GetURL(ctx context.Context, shortCode string) (string, error)
	DeleteURL(ctx context.Context, shortCode string) error

	// Rate Limiting

	IncrementRateLimit(ctx context.Context, key string, window time.Duration) (int64, error)
	GetRateLimit(ctx context.Context, key string) (int64, error)

	// Analytics Caching

	IncrementClickCount(ctx context.Context, urlID int64) error
	GetClickCount(ctx context.Context, urlID int64) (int64, error)

	// User Session Caching

	SetUserSession(ctx context.Context, token string, user interface{}, ttl time.Duration) error
	GetUserSession(ctx context.Context, token string) (interface{}, error)
	DeleteUserSession(ctx context.Context, token string) error

	// Generic Cache Operations

	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type cacheService struct {
	redisClient *database.RedisClient
	memCache    map[string]interface{}
}

func NewCacheService(redisClient *database.RedisClient) CacheService {
	return &cacheService{
		redisClient: redisClient,
		memCache:    map[string]interface{}{},
	}
}

func (s *cacheService) SetURL(ctx context.Context, shortCode, originalURL string, ttl time.Duration) error {
	s.memCache[shortCode] = originalURL
	err := s.redisClient.Set(ctx, shortCode, originalURL, ttl)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacheService) GetURL(ctx context.Context, shortCode string) (string, error) {
	value := s.memCache[shortCode]
	if value != nil {
		return value.(string), nil
	}
	value, err := s.Get(ctx, shortCode)
	if err != nil {
		return "", err
	}
	return value.(string), nil
}

func (s *cacheService) DeleteURL(ctx context.Context, shortCode string) error {
	s.memCache[shortCode] = nil
	err := s.redisClient.Del(ctx, shortCode)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacheService) IncrementRateLimit(ctx context.Context, key string, window time.Duration) (int64, error) {
	// TODO: Increment counter with automatic expiration

	limit, err := s.redisClient.Incr(ctx, key)
	if err != nil {
		return 0, err
	}
	if limit == 1 {

	}

	return 0, nil
}

func (s *cacheService) GetRateLimit(ctx context.Context, key string) (int64, error) {
	value := s.memCache[key]
	if value != nil {
		limit, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return limit, nil
	} else {
		value, err := s.redisClient.Get(ctx, key)
		if err != nil {
			return 0, err
		}
		limit, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, err
		}
		return limit, nil
	}
}

func (s *cacheService) IncrementClickCount(ctx context.Context, urlID int64) error {
	_, err := s.redisClient.Incr(ctx, strconv.FormatInt(urlID, 10))
	if err != nil {
		return err
	}
	return nil
}

func (s *cacheService) GetClickCount(ctx context.Context, urlID int64) (int64, error) {
	value, err := s.redisClient.Get(ctx, strconv.FormatInt(urlID, 10))
	if err != nil {
		return 0, err
	}
	clickCount, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return clickCount, nil
}

func (s *cacheService) SetUserSession(ctx context.Context, token string, user interface{}, ttl time.Duration) error {
	err := s.redisClient.Set(ctx, token, user, ttl)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacheService) GetUserSession(ctx context.Context, token string) (interface{}, error) {
	return s.redisClient.Get(ctx, token)
}

func (s *cacheService) DeleteUserSession(ctx context.Context, token string) error {
	return s.redisClient.Del(ctx, token)
}

func (s *cacheService) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := s.redisClient.Set(ctx, key, value, ttl)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacheService) Get(ctx context.Context, key string) (interface{}, error) {
	return s.redisClient.Get(ctx, key)
}

func (s *cacheService) Delete(ctx context.Context, key string) error {
	return s.redisClient.Del(ctx, key)
}

func (s *cacheService) Exists(ctx context.Context, key string) (bool, error) {
	return s.redisClient.Exists(ctx, key)
}
