package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vickykumar/url_shortner/internal/models"
	"github.com/vickykumar/url_shortner/internal/service"
	"net/http"
	"strconv"
)

type URLHandler struct {
	urlService       service.URLService
	analyticsService service.AnalyticsService
}

func NewURLHandler(urlService service.URLService, analyticsService service.AnalyticsService) *URLHandler {
	return &URLHandler{
		urlService:       urlService,
		analyticsService: analyticsService,
	}
}

func (h *URLHandler) CreateURL(c *gin.Context) {
	var userId = c.Param("user_id")
	var req models.CreateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.ValidateCreateUrl(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	resp, err := h.urlService.CreateURL(c, &userIdInt, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (h *URLHandler) GetURL(c *gin.Context) {
	// TODO: Parse URL parameter (ID)
	// TODO: Get user from context
	// TODO: Call service to get URL
	// TODO: Return response
	
}

func (h *URLHandler) UpdateURL(c *gin.Context) {
	// TODO: Parse URL parameter and request body
	// TODO: Validate input
	// TODO: Get user from context
	// TODO: Call service to update URL
	// TODO: Return response
}

func (h *URLHandler) DeleteURL(c *gin.Context) {
	// TODO: Parse URL parameter (ID)
	// TODO: Get user from context
	// TODO: Call service to delete URL
	// TODO: Return response
}

func (h *URLHandler) ListURLs(c *gin.Context) {
	// TODO: Parse query parameters
	// TODO: Get user from context
	// TODO: Call service to list URLs
	// TODO: Return paginated response
}

func (h *URLHandler) RedirectURL(c *gin.Context) {
	// TODO: Parse short code from URL
	// TODO: Extract analytics data from request
	// TODO: Call service to get original URL
	// TODO: Check for password protection
	// TODO: Redirect to original URL
}

func (h *URLHandler) PreviewURL(c *gin.Context) {
	// TODO: Parse short code from URL
	// TODO: Get URL details without redirecting
	// TODO: Return URL preview information
}
