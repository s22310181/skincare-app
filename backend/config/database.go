package config

import (
	"fmt"
	"log"
	"os"

	"skincare-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Load .env
	_ = godotenv.Load()

	// Ambil dari environment
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Default config (AMAN XAMPP)
	if dbUser == "" {
		dbUser = "root"
	}

	// 🔥 FIX UTAMA DI SINI
	// kalau kosong ATAU masih "password", paksa jadi kosong
	if dbPassword == "" || dbPassword == "password" {
		dbPassword = ""
	}

	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	if dbPort == "" {
		dbPort = "3306"
	}

	if dbName == "" {
		dbName = "skincare_db"
	}

	// Debug log (biar kamu lihat config yang dipakai)
	log.Printf("DB CONFIG → user:%s host:%s port:%s db:%s",
		dbUser, dbHost, dbPort, dbName)

	// DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Koneksi
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}

	log.Println("✅ Database berhasil terkoneksi")

	// Auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.SkinAnalysis{},
	)
	if err != nil {
		log.Fatal("❌ Gagal migrate:", err)
	}

	log.Println("✅ Migration sukses")

	return db
}