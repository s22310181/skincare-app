package controllers

import (
	"net/http"
	"os"
	"skincare-app/middleware"
	"skincare-app/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Register handles user registration
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	existingUser, _ := models.GetUserByEmail(db, req.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := models.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate token
	token := generateToken(user.ID, user.Email)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
		"user": UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// Login handles user login
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user
	user, err := models.GetUserByEmail(db, req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate token
	token := generateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// GetProfile gets user profile
func GetProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := middleware.GetUserID(c)

	user, err := models.GetUserByID(db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func generateToken(userID uint, email string) string {
	claims := &middleware.Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString
}
