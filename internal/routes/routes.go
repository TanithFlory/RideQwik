package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rideqwik/api/internal/handlers"
	"github.com/rideqwik/api/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler())

	router.GET("/health", handlers.HealthCheck)
}
