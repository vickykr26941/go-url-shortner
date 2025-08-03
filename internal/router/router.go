package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vickykumar/url_shortner/internal/handler"
	"net/http"
)

type Router interface {
	Start(ctx context.Context) error
}

type impl struct {
	engine      *gin.Engine
	authHandler *handler.AuthHandler
}

func NewRouter(authHandler *handler.AuthHandler) (Router, error) {
	return &impl{
		engine:      gin.New(),
		authHandler: authHandler,
	}, nil
}

func (r *impl) Start(ctx context.Context) error {
	r.engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// make a group later for user, urls, and other handlers as required.

	// auth and user APIs
	r.engine.POST("/api/v1/register/user", r.authHandler.Register)
	r.engine.POST("/api/v1/login", r.authHandler.Login)
	r.engine.POST("/api/v1/logout", r.authHandler.Logout)
	r.engine.GET("/api/v1/user/profile/:userId", r.authHandler.GetProfile)
	r.engine.POST("/api/v1/user/profile/:userId", r.authHandler.UpdateProfile)
	r.engine.POST("/api/v1/refresh/token", r.authHandler.RefreshToken)
	r.engine.GET("/api/v1/key/generate", r.authHandler.GenerateAPIKey)
	r.engine.GET("/api/v1/key/revoke", r.authHandler.RevokeAPIKey)

	// url apis
	return r.engine.Run("0.0.0.0:80")
}
