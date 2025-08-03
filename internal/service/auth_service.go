package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vickykumar/url_shortner/internal/config"
	"github.com/vickykumar/url_shortner/internal/models"
	"github.com/vickykumar/url_shortner/internal/repository"
	"github.com/vickykumar/url_shortner/pkg"
	"strconv"
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

	StoreRefreshToken(ctx context.Context, userID int64, refreshToken, jti string) error
}

type authService struct {
	userRepo     repository.UserRepository
	cacheService CacheService
	tokenExpiry  time.Duration
	authConfig   config.AuthConfig
}

func NewAuthService(
	userRepo repository.UserRepository,
	cacheService CacheService,
	authConfig config.AuthConfig,
) AuthService {
	return &authService{
		userRepo:     userRepo,
		cacheService: cacheService,
		authConfig:   authConfig,
	}
}

func (s *authService) Register(ctx context.Context, req *models.RegisterRequest) (*models.LoginResponse, error) {
	userId := time.Now().UTC().UnixNano()
	user := &models.User{
		Email:           req.Email,
		PasswordHash:    req.Password, // skipping hashing for now, should be hashed
		Name:            &req.Name,
		IsPremium:       false,
		DailyClickCount: models.MAX_DAILY_CLICK_COUNT,
		DailyURLCount:   models.MAX_DAILY_URL_COUNT,
		ID:              userId,
	}

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	jti := uuid.New().String()
	accessToken, refreshToken, err := s.generateJWT(ctx, userId, jti)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	loginResponse := &models.LoginResponse{
		User: models.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			IsPremium: user.IsPremium,
			CreatedAt: user.CreatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	if err := s.userRepo.UpdateLastLogin(ctx, userId); err != nil {
		fmt.Printf("Failed to update last login for user %d: %v\n", userId, err)
	}
	if err := s.StoreRefreshToken(ctx, userId, refreshToken, jti); err != nil {
		return nil, fmt.Errorf("failed to store refresh token in cache: %w", err)
	}
	return loginResponse, nil
}

func (s *authService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user.PasswordHash != req.Password {
		return nil, fmt.Errorf("invalid credentials")
	}

	jti := uuid.New().String()
	accessToken, refreshToken, err := s.generateJWT(ctx, user.ID, jti)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	loginResponse := &models.LoginResponse{
		User: models.UserResponse{
			ID:          user.ID,
			Email:       user.Email,
			Name:        user.Name,
			IsPremium:   user.IsPremium,
			CreatedAt:   user.CreatedAt,
			LastLoginAt: user.LastLoginAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = s.userRepo.UpdateLastLogin(ctx, user.ID)
	if err != nil {
		fmt.Printf("Failed to update last login for user %d: %v\n", user.ID, err)
	}

	if err := s.StoreRefreshToken(ctx, user.ID, refreshToken, jti); err != nil {
		return nil, fmt.Errorf("failed to store refresh token in cache: %w", err)
	}
	return loginResponse, nil
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

func (s *authService) generateJWT(ctx context.Context, userID int64, jti string) (accessToken, refreshToken string, err error) {
	user, err := s.userRepo.GetByID(context.Background(), userID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user by ID: %w", err)
	}

	jwtClaims := &pkg.UserToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.authConfig.RefreshTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "vicky.url.shortner",
			Subject:   fmt.Sprintf("%d", userID),
		},
		UserId: strconv.FormatInt(userID, 10),
		Email:  user.Email,
	}

	jwtClaimMap := jwt.MapClaims{}

	jwtClaimMap["iss"], _ = jwtClaims.GetIssuer()
	jwtClaimMap["sub"], _ = jwtClaims.GetSubject()
	jwtClaimMap["exp"], _ = jwtClaims.GetExpirationTime()
	jwtClaimMap["iat"], _ = jwtClaims.GetIssuedAt()
	jwtClaimMap["user_id"] = jwtClaims.UserId
	jwtClaimMap["jti"] = jwtClaims.ID

	if user.Email != "" {
		jwtClaimMap["email"] = user.Email
	}

	if user.DailyClickCount != 0 {
		jwtClaimMap["daily_c_cnt"] = user.DailyClickCount
	}

	if user.DailyURLCount != 0 {
		jwtClaimMap["daily_u_cnt"] = user.DailyURLCount
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaimMap).SignedString([]byte(s.authConfig.JWTSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken = uuid.New().String()
	err = s.StoreRefreshToken(ctx, userID, refreshToken, jti)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) StoreRefreshToken(ctx context.Context, userID int64, refreshToken, jti string) error {
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user by ID: %w", err)
	}
	values := map[string]interface{}{
		"jti":           jti,
		"user_id":       userID,
		"refresh_token": refreshToken,
	}
	userRefreshKey := fmt.Sprintf("user:%d:refresh_token", userID)

	ttl := s.authConfig.RefreshTokenExpiry
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("failed to marshal refresh token values: %w", err)
	}
	if err := s.cacheService.Set(ctx, userRefreshKey, jsonValue, ttl); err != nil {
		return fmt.Errorf("failed to store refresh token in cache: %w", err)
	}
	return nil
}
