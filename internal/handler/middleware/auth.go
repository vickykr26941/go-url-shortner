package middleware

import (
	"context"
	"github.com/vickykumar/url_shortner/internal/service"
	"net/http"
)

type AuthMiddleware struct {
	authService service.AuthService
}

func NewAuthMiddleware(authService service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) RequireAuth() func(http.Handler) http.Handler {
	// TODO: Middleware to require authentication
	// TODO: Extract token from Authorization header
	// TODO: Validate token and get user
	// TODO: Set user in request context
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Implementation
			next.ServeHTTP(w, r)
		})
	}
}

func (m *AuthMiddleware) OptionalAuth() func(http.Handler) http.Handler {
	// TODO: Middleware for optional authentication
	// TODO: Set user in context if token is valid
	// TODO: Continue even if no token or invalid token
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Implementation
			next.ServeHTTP(w, r)
		})
	}
}

func (m *AuthMiddleware) RequireAPIKey() func(http.Handler) http.Handler {
	// TODO: Middleware to require API key authentication
	// TODO: Extract API key from header or query parameter
	// TODO: Validate API key and get user
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Implementation
			next.ServeHTTP(w, r)
		})
	}
}

func (m *AuthMiddleware) getUserFromToken(token string) (interface{}, error) {
	// TODO: Helper function to get user from JWT token
	return nil, nil
}

func (m *AuthMiddleware) getUserFromAPIKey(apiKey string) (interface{}, error) {
	// TODO: Helper function to get user from API key
	return nil, nil
}

func (m *AuthMiddleware) setUserInContext(ctx context.Context, user interface{}) context.Context {
	// TODO: Helper function to set user in request context
	return ctx
}
