package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"-" gorm:"not null"`
	Analysis []SkinAnalysis `json:"analysis,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

// Get user by email
func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by ID
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create user
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
