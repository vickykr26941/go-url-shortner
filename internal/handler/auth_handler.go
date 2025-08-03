package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vickykumar/url_shortner/internal/models"
	"github.com/vickykumar/url_shortner/internal/service"
	"net/http"
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
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := req.ValidateCreateRequest(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	resp, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := req.ValidateLoginRequest(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	resp, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	resp, err := h.authService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
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
