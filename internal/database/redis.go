package database

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vickykumar/url_shortner/internal/config"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisConnection(config *config.RedisConfig) (*RedisClient, error) {

	port, _ := strconv.ParseInt(config.Port, 10, 64)
	opts := &redis.Options{
		Addr:            fmt.Sprintf("%s:%d", config.Host, port),
		Password:        config.Password,
		DB:              config.DB,
		PoolSize:        config.PoolSize,
		MinIdleConns:    int(config.IdleTimeout),
		MaxIdleConns:    int(config.IdleTimeout),
		ConnMaxIdleTime: time.Duration(config.IdleTimeout) * time.Second,
		ConnMaxLifetime: time.Duration(config.IdleTimeout) * time.Second,
		DialTimeout:     time.Duration(config.IdleTimeout) * time.Second,
		ReadTimeout:     time.Duration(config.IdleTimeout) * time.Second,
		WriteTimeout:    time.Duration(config.IdleTimeout) * time.Second,
	}

	rdb := redis.NewClient(opts)
	redisClient := &RedisClient{client: rdb}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx); err != nil {
		rdb.Close()
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}

	return redisClient, nil
}

func (r *RedisClient) Close() error {
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

func (r *RedisClient) Ping(ctx context.Context) error {
	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return r.client.Ping(ctx).Err()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("redis client is nil")
	}

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", fmt.Errorf("key '%s' not found", key)
		}
		return "", err
	}

	return result, nil
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}

	if len(keys) == 0 {
		return nil // Nothing to delete
	}

	return r.client.Del(ctx, keys...).Err()
}

func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	if r.client == nil {
		return false, fmt.Errorf("redis client is nil")
	}

	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func (r *RedisClient) Incr(ctx context.Context, key string) (int64, error) {
	if r.client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}

	result, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}

	return r.client.Expire(ctx, key, expiration).Err()
}

// Additional helper methods for common operations

func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	if r.client == nil {
		return false, fmt.Errorf("redis client is nil")
	}

	return r.client.SetNX(ctx, key, value, expiration).Result()
}

func (r *RedisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	if r.client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}

	return r.client.TTL(ctx, key).Result()
}

func (r *RedisClient) GetClient() *redis.Client {
	return r.client
}
