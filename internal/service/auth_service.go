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
	"github.com/vickykumar/url_shortner/internal/utils"
	"github.com/vickykumar/url_shortner/pkg"
	"strconv"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.LoginResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
	RefreshToken(ctx context.Context, req *models.TokenData) (*models.LoginResponse, error)
	GenerateAPIKey(ctx context.Context, userID int64) (*models.APIKeyResponse, error)
	RevokeAPIKey(ctx context.Context, userID int64) error
	ChangePassword(ctx context.Context, userID int64, request *models.UpdatePassRequest) error

	StoreRefreshToken(ctx context.Context, userID int64, refreshToken, jti string) error
	Logout(ctx context.Context, req *models.TokenData) error

	GetUserProfile(ctx context.Context, userId int64) (*models.UserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *models.UpdateUserRequest) error

	ValidateToken(ctx context.Context, req *models.TokenData) (*pkg.UserToken, error)
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
	accessToken, refreshToken, err := s.GenerateJwtToken(ctx, userId, jti)
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
	accessToken, refreshToken, err := s.GenerateJwtToken(ctx, user.ID, jti)
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

func (s *authService) ValidateToken(ctx context.Context, req *models.TokenData) (*pkg.UserToken, error) {
	refreshData, err := s.cacheService.Get(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	var refreshInfo models.RefreshInfo
	if refreshDataBytes, ok := refreshData.([]byte); ok {
		if err := json.Unmarshal(refreshDataBytes, &refreshInfo); err != nil {
			return nil, err
		}
	}

	refreshToken := refreshInfo.RefreshToken
	accessToken := req.AccessToken
	if refreshToken == "" || accessToken == "" {
		return nil, fmt.Errorf("access token & refresh token required to refresh")
	}
	token, err := s.ParseJwtToken(accessToken)
	if err != nil {
		return nil, err
	}

	switch {
	case token.UserId != refreshInfo.UserId:
		return nil, fmt.Errorf("invalid refresh token")
	case token.ID != refreshInfo.Jti:
		return nil, fmt.Errorf("invalid refresh token")
	}

	return token, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *models.TokenData) (*models.LoginResponse, error) {

	token, err := s.ValidateToken(ctx, req)
	if err != nil {
		return nil, err
	}
	userId, _ := strconv.ParseInt(token.UserId, 10, 64)
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	jti := uuid.New().String()
	accessToken, refreshToken, err := s.GenerateJwtToken(ctx, user.ID, jti)
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

	return loginResponse, nil
}

func (s *authService) GenerateAPIKey(ctx context.Context, userID int64) (*models.APIKeyResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	apiKey, err := utils.GenerateAPIKey()
	if err != nil {
		return nil, err
	}

	user.APIKey = &apiKey
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &models.APIKeyResponse{
		APIKey:    apiKey,
		CreatedAt: time.Now(),
	}
	return response, nil
}

func (s *authService) RevokeAPIKey(ctx context.Context, userID int64) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	user.APIKey = nil
	err = s.userRepo.Update(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ChangePassword(ctx context.Context, userID int64, updatePassReq *models.UpdatePassRequest) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.PasswordHash == *updatePassReq.NewPassword {
		return fmt.Errorf("new password is the same as the old password")
	}

	user.PasswordHash = *updatePassReq.NewPassword
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) GenerateJwtToken(ctx context.Context, userID int64, jti string) (accessToken, refreshToken string, err error) {
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

func (s *authService) ParseJwtToken(token string) (*pkg.UserToken, error) {
	userToken, err := jwt.ParseWithClaims(token, &pkg.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.authConfig.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := userToken.Claims.(*pkg.UserToken); ok && userToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
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

	// userRefreshKey := fmt.Sprintf("user:%d:refresh_token", userID)
	ttl := s.authConfig.RefreshTokenExpiry
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("failed to marshal refresh token values: %w", err)
	}
	if err := s.cacheService.Set(ctx, refreshToken, jsonValue, ttl); err != nil {
		return fmt.Errorf("failed to store refresh token in cache: %w", err)
	}
	return nil
}

func (s *authService) Logout(ctx context.Context, req *models.TokenData) error {
	err := s.cacheService.Delete(ctx, req.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to delete refresh token from cache: %w", err)
	}
	return nil
}

func (s *authService) GetUserProfile(ctx context.Context, userId int64) (*models.UserProfileResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	response := &models.UserProfileResponse{
		ID:              user.ID,
		Email:           user.Email,
		Name:            user.Name,
		APIKey:          user.APIKey,
		IsPremium:       user.IsPremium,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		LastLoginAt:     user.LastLoginAt,
		DailyURLCount:   user.DailyURLCount,
		DailyClickCount: user.DailyClickCount,
		LastResetDate:   user.LastResetDate,
	}

	return response, nil
}

func (s *authService) UpdateUserProfile(ctx context.Context, request *models.UpdateUserRequest) error {
	user, err := s.userRepo.GetByEmail(ctx, *request.Email)
	if err != nil {
		return fmt.Errorf("failed to get user by email: %w", err)
	}

	if user.PasswordHash != *request.Password {
		return fmt.Errorf("password does not match")
	}

	user.Name = request.Name
	user.Email = *request.Email
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
