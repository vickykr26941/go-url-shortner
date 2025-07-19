package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vickykumar/url_shortner/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// TODO: Parse registration request
	// TODO: Validate input
	// TODO: Call service to register user
	// TODO: Return success response with tokens
}

func (h *AuthHandler) Login(c *gin.Context) {
	// TODO: Parse login request
	// TODO: Validate input
	// TODO: Call service to authenticate user
	// TODO: Return success response with tokens
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: Parse refresh token request
	// TODO: Call service to refresh token
	// TODO: Return new tokens
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// TODO: Get user from context
	// TODO: Invalidate user session
	// TODO: Return success response
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	// TODO: Get user from context
	// TODO: Return user profile data
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	// TODO: Parse update request
	// TODO: Get user from context
	// TODO: Update user profile
	// TODO: Return updated profile
}

func (h *AuthHandler) GenerateAPIKey(c *gin.Context) {
	// TODO: Get user from context
	// TODO: Generate new API key
	// TODO: Return API key response
}

func (h *AuthHandler) RevokeAPIKey(c *gin.Context) {
	// TODO: Get user from context
	// TODO: Revoke API key
	// TODO: Return success response
}
