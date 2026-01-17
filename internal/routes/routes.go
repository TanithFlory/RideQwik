package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rideqwik/api/internal/config"
	"github.com/rideqwik/api/internal/handlers"
	"github.com/rideqwik/api/internal/middleware"
	"github.com/rideqwik/api/internal/repositories"
	"github.com/rideqwik/api/internal/services"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, pool *pgxpool.Pool) {
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler())

	router.GET("/health", handlers.HealthCheck)

	userRepo := repositories.NewUserRepository(pool)

	authService := services.NewAuthService(userRepo, cfg)
	uberService := services.NewUberService(cfg)
	rideService := services.NewRideService(uberService)

	authHandler := handlers.NewAuthHandler(authService)
	rideHandler := handlers.NewRideRequestHandler(rideService)
	uberOAuthHandler := handlers.NewUberOAuthHandler(cfg, uberService)

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		oauth := v1.Group("/oauth")
		oauth.Use(middleware.AuthMiddleware(authService))
		{
			oauth.GET("/uber/authorize", uberOAuthHandler.AuthorizeUber)
		}

		v1.GET("/oauth/uber/callback", uberOAuthHandler.UberCallback)

		rides := v1.Group("/rides")
		rides.Use(middleware.AuthMiddleware(authService))
		{
			rides.POST("/requests", rideHandler.RequestRides)
		}

		user := v1.Group("/user")
		user.Use(middleware.AuthMiddleware(authService))
		{
			user.GET("/me", authHandler.GetMe)
		}
	}
}
