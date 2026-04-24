package routes

import (
	"skincare-app/controllers"
	"skincare-app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Middleware to pass DB to handlers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Public routes
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Protected routes
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		// User profile
		protectedRoutes.GET("/profile", controllers.GetProfile)

		// Skin analysis
		protectedRoutes.POST("/skin-analysis", controllers.AnalyzeSkin)
		protectedRoutes.GET("/history", controllers.GetHistory)
		protectedRoutes.GET("/history/:id", controllers.GetAnalysisDetail)
	}
}
