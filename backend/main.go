package main

import (
	"log"
	"os"
	"skincare-app/config"
	"skincare-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database (FIX DI SINI)
	db := config.InitDB()

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}

	// Create router
	router := gin.Default()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Serve static files
	router.Static("/assets", "../frontend/assets")
	router.StaticFile("/", "../frontend/index.html")
	router.StaticFile("/login", "../frontend/login.html")
	router.StaticFile("/register", "../frontend/register.html")
	router.StaticFile("/dashboard", "../frontend/dashboard.html")
	router.StaticFile("/analysis", "../frontend/analysis.html")

	// Initialize routes
	routes.SetupRoutes(router, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}