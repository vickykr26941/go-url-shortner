package service

import (
	"context"
	"github.com/vickykumar/url_shortner/internal/models"
	"github.com/vickykumar/url_shortner/internal/repository"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.LoginResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
	RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.LoginResponse, error)
	ValidateToken(ctx context.Context, token string) (*models.User, error)
	ValidateAPIKey(ctx context.Context, apiKey string) (*models.User, error)
	GenerateAPIKey(ctx context.Context, userID int64) (*models.APIKeyResponse, error)
	RevokeAPIKey(ctx context.Context, userID int64) error
	ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error
}

type authService struct {
	userRepo     repository.UserRepository
	cacheService CacheService
	jwtSecret    string
	tokenExpiry  time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	cacheService CacheService,
	jwtSecret string,
	tokenExpiry time.Duration,
) AuthService {
	return &authService{
		userRepo:     userRepo,
		cacheService: cacheService,
		jwtSecret:    jwtSecret,
		tokenExpiry:  tokenExpiry,
	}
}

func (s *authService) Register(ctx context.Context, req *models.RegisterRequest) (*models.LoginResponse, error) {
	// TODO: Validate email uniqueness
	// TODO: Hash password
	// TODO: Create user in database
	// TODO: Generate JWT tokens
	// TODO: Update last login
	return nil, nil
}

func (s *authService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	// TODO: Get user by email
	// TODO: Validate password
	// TODO: Generate JWT tokens
	// TODO: Update last login timestamp
	// TODO: Cache user session
	return nil, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.LoginResponse, error) {
	// TODO: Validate refresh token
	// TODO: Generate new access token
	// TODO: Optionally generate new refresh token
	return nil, nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*models.User, error) {
	// TODO: Parse and validate JWT token
	// TODO: Check token expiration
	// TODO: Get user from cache or database
	return nil, nil
}

func (s *authService) ValidateAPIKey(ctx context.Context, apiKey string) (*models.User, error) {
	// TODO: Get user by API key from cache first
	// TODO: Fallback to database if not cached
	// TODO: Cache user data
	return nil, nil
}

func (s *authService) GenerateAPIKey(ctx context.Context, userID int64) (*models.APIKeyResponse, error) {
	// TODO: Generate secure random API key
	// TODO: Update user record with new API key
	// TODO: Invalidate cache
	return nil, nil
}

func (s *authService) RevokeAPIKey(ctx context.Context, userID int64) error {
	// TODO: Remove API key from user record
	// TODO: Invalidate cache
	return nil
}

func (s *authService) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	// TODO: Get user and validate old password
	// TODO: Hash new password
	// TODO: Update user record
	// TODO: Invalidate all user sessions
	return nil
}

func (s *authService) hashPassword(password string) (string, error) {
	// TODO: Hash password using bcrypt
	return "", nil
}

func (s *authService) validatePassword(password, hash string) bool {
	// TODO: Validate password against bcrypt hash
	return false
}

func (s *authService) generateJWT(userID int64) (accessToken, refreshToken string, err error) {
	// TODO: Generate JWT access token
	// TODO: Generate JWT refresh token
	return "", "", nil
}
