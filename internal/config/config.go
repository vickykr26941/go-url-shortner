package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Auth     AuthConfig     `yaml:"auth"`
	Cache    CacheConfig    `yaml:"cache"`
	Logger   LoggerConfig   `yaml:"logger"`
	Metrics  MetricsConfig  `yaml:"metrics"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Host            string        `yaml:"host"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type DatabaseConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	User         string        `yaml:"user"`
	Password     string        `yaml:"password"`
	Database     string        `yaml:"database"`
	SSLMode      string        `yaml:"ssl_mode"`
	MaxOpenConns int           `yaml:"max_open_conns"`
	MaxIdleConns int           `yaml:"max_idle_conns"`
	MaxLifetime  time.Duration `yaml:"max_lifetime"`
}

type RedisConfig struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	Password    string        `yaml:"password"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	PoolSize    int           `yaml:"pool_size"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type AuthConfig struct {
	JWTSecret          string        `yaml:"jwt_secret"`
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
	PasswordMinLength  int           `yaml:"password_min_length"`
	BCryptCost         int           `yaml:"bcrypt_cost"`
	ServiceUrl         string        `yaml:"service_url"`
}

type CacheConfig struct {
	URLCacheTTL       time.Duration `yaml:"url_cache_ttl"`
	AnalyticsTTL      time.Duration `yaml:"analytics_ttl"`
	RateLimitTTL      time.Duration `yaml:"rate_limit_ttl"`
	InMemoryCacheSize int           `yaml:"inmemory_cache_size"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Output string `yaml:"output"`
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
	Port    string `yaml:"port"`
}

func Load(envKey string) (*Config, error) {
	jsonFromEnv := os.Getenv(envKey)
	ymlFile := false
	var jsonFromYml []byte
	var err error
	if jsonFromEnv == "" {
		jsonFromYml, err = os.ReadFile("internal/config/config.yml")
		if err != nil {
			return nil, err
		}
		ymlFile = true
	}
	var config Config
	switch ymlFile {
	case true:
		err := yaml.Unmarshal(jsonFromYml, &config)
		if err != nil {
			return nil, err
		}
	default:
		err := json.Unmarshal([]byte(jsonFromEnv), &config)
		if err != nil {
			return nil, err
		}
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) Validate() error {
	// validate and set default values for configurations
	if c.Server.Port == "" {
		c.Server.Port = "8080"
	}
	if c.Server.Host == "" {
		c.Server.Host = "127.0.0.1"
	}
	if c.Server.ShutdownTimeout == 0 {
		c.Server.ShutdownTimeout = 5 * time.Second
	}

	return nil
}
