package server

import (
	config2 "github.com/vickykumar/url_shortner/internal/config"
	"github.com/vickykumar/url_shortner/internal/database"
	"github.com/vickykumar/url_shortner/internal/handler"
	"github.com/vickykumar/url_shortner/internal/repository"
	router2 "github.com/vickykumar/url_shortner/internal/router"
	"github.com/vickykumar/url_shortner/internal/service"
)

// TODO: use google wire for better dependencies injection
func initializerRoute() (router2.Router, error) {
	config, err := config2.Load("CONFIG_JSON")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	mysql, err := database.NewMySqlConnection(&config.Database)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	redis, err := database.NewRedisConnection(&config.Redis)
	if err != nil {
		panic("Failed to connect to redis: " + err.Error())
	}
	userRepo := repository.NewUserRepository(mysql.GetDB())
	cacheService := service.NewCacheService(redis)
	userService := service.NewAuthService(userRepo, cacheService, config.Auth)
	authHandler := handler.NewAuthHandler(userService)

	router, err := router2.NewRouter(authHandler)
	if err != nil {
		panic("Failed to create router: " + err.Error())
	}
	return router, err
}
