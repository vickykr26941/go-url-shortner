package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/vickykumar/url_shortner/internal/config"
	"github.com/vickykumar/url_shortner/internal/models"
	"github.com/vickykumar/url_shortner/internal/repository"
	"time"
)

type URLService interface {
	CreateURL(ctx context.Context, userID *int64, req *models.CreateURLRequest) (*models.URLResponse, error)
	GetURL(ctx context.Context, id int64, userID *int64) (*models.URLResponse, error)
	UpdateURL(ctx context.Context, id int64, userID int64, req *models.UpdateURLRequest) (*models.URLResponse, error)
	DeleteURL(ctx context.Context, id int64, userID int64) error
	ListURLs(ctx context.Context, userID int64, req *models.URLListRequest) ([]*models.URLResponse, int64, error)
	RedirectURL(ctx context.Context, shortCode string, req *models.AnalyticsRequest) (string, error)
	ValidateURLPassword(ctx context.Context, shortCode string, password string) error
}

type urlService struct {
	urlRepo       repository.URLRepository
	analyticsRepo repository.AnalyticsRepository
	tagRepo       repository.TagRepository
	cacheService  CacheService
	authConfig    config.AuthConfig
}

func NewURLService(
	urlRepo repository.URLRepository,
	analyticsRepo repository.AnalyticsRepository,
	tagRepo repository.TagRepository,
	cacheService CacheService,
	authConfig config.AuthConfig,
) URLService {
	return &urlService{
		urlRepo:       urlRepo,
		analyticsRepo: analyticsRepo,
		tagRepo:       tagRepo,
		cacheService:  cacheService,
		authConfig:    authConfig,
	}
}

func (s *urlService) CreateURL(ctx context.Context, userID *int64, req *models.CreateURLRequest) (*models.URLResponse, error) {
	if req.OriginalURL == "" {
		return nil, errors.New("original URL is required")
	}

	shortCode, err := s.generateShortCode(req.CustomCode)
	if err != nil {
		return nil, err
	}

	var passwordHash *string
	if req.Password != nil {
		//hash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		//if err != nil {
		//	return nil, fmt.Errorf("failed to hash password: %w", err)
		//}

		hash := *req.Password
		str := string(hash)
		passwordHash = &str
	}

	url := &models.URL{
		ShortCode:    shortCode,
		OriginalURL:  req.OriginalURL,
		UserID:       userID,
		Title:        req.Title,
		Description:  req.Description,
		ExpiresAt:    req.ExpiresAt,
		IsActive:     true,
		ClickCount:   0,
		IsCustom:     req.CustomCode != nil,
		PasswordHash: passwordHash,
	}

	if err := s.urlRepo.Create(ctx, url); err != nil {
		return nil, err
	}

	return s.urlSuccessResponse(url)
}

func (s *urlService) GetURL(ctx context.Context, id int64, userID *int64) (*models.URLResponse, error) {
	url, err := s.urlRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userID != nil && (url.UserID == nil || *url.UserID != *userID) {
		return nil, errors.New("unauthorized access")
	}

	return s.urlSuccessResponse(url)
}

func (s *urlService) GetURLByShortCode(ctx context.Context, shortCode string) (*models.URLResponse, error) {
	url, err := s.urlRepo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return s.urlSuccessResponse(url)
}

func (s *urlService) UpdateURL(ctx context.Context, id int64, userID int64, req *models.UpdateURLRequest) (*models.URLResponse, error) {
	url, err := s.urlRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if url.UserID == nil || *url.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	url.Title = req.Title
	url.Description = req.Description
	url.ExpiresAt = req.ExpiresAt
	url.IsActive = *req.IsActive

	if err := s.urlRepo.Update(ctx, url); err != nil {
		return nil, err
	}
	return s.urlSuccessResponse(url)
}

func (s *urlService) DeleteURL(ctx context.Context, id int64, userID int64) error {
	url, err := s.urlRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if url.UserID == nil || *url.UserID != userID {
		return errors.New("unauthorized")
	}
	return s.urlRepo.Delete(ctx, id)
}

func (s *urlService) ListURLs(ctx context.Context, userID int64, req *models.URLListRequest) ([]*models.URLResponse, int64, error) {
	urls, count, err := s.urlRepo.List(ctx, userID, req)
	if err != nil {
		return nil, 0, err
	}
	var responses []*models.URLResponse
	for _, u := range urls {
		resp, _ := s.urlSuccessResponse(u)
		responses = append(responses, resp)
	}
	return responses, count, nil
}

func (s *urlService) RedirectURL(ctx context.Context, shortCode string, req *models.AnalyticsRequest) (string, error) {
	url, err := s.urlRepo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}
	if !url.IsActive || url.ExpiresAt.Before(time.Now()) {
		return "", errors.New("url is expired or inactive")
	}
	go func() {
		err := s.urlRepo.IncrementClickCount(context.Background(), url.ID)
		if err != nil {
			fmt.Println(err)
		}
	}()
	// optionally record analytics via s.analyticsRepo here
	return url.OriginalURL, nil
}

func (s *urlService) ValidateURLPassword(ctx context.Context, shortCode string, password string) error {
	url, err := s.urlRepo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return err
	}
	if url.PasswordHash == nil {
		return nil
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(*url.PasswordHash), []byte(password)); err != nil {
	//	return errors.New("invalid password")
	//}
	return nil
}

func (s *urlService) generateShortCode(customCode *string) (string, error) {
	if customCode != nil {
		exists, err := s.urlRepo.GetByShortCode(context.Background(), *customCode)
		if err == nil && exists != nil {
			return "", errors.New("custom short code already exists")
		}
		return *customCode, nil
	}

	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%x", timestamp), nil
}

func (s *urlService) urlSuccessResponse(url *models.URL) (*models.URLResponse, error) {

	response := &models.URLResponse{
		ID:          url.ID,
		ShortCode:   url.ShortCode,
		ShortURL:    s.authConfig.ServiceUrl + "/" + url.ShortCode,
		OriginalURL: url.OriginalURL,
		Title:       url.Title,
		Description: url.Description,
		ExpiresAt:   url.ExpiresAt,
		ClickCount:  url.ClickCount,
		IsCustom:    url.IsCustom,
	}

	return response, nil
}
