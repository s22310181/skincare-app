package controllers

import (
	"net/http"
	"skincare-app/middleware"
	"skincare-app/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SkinAnalysisRequest struct {
	SkinType    string   `json:"skin_type" binding:"required"`     // kering, berminyak, kombinasi, sensitif
	SkinIssues  []string `json:"skin_issues" binding:"required"`   // jerawat, kusam, flek hitam, dll
	Sensitivity string   `json:"sensitivity" binding:"required"`   // rendah, sedang, tinggi
}

type SkinAnalysisResponse struct {
	ID             uint   `json:"id"`
	SkinType       string `json:"skin_type"`
	SkinIssues     []string `json:"skin_issues"`
	Sensitivity    string `json:"sensitivity"`
	MorningRoutine models.RoutineData `json:"morning_routine"`
	NightRoutine   models.RoutineData `json:"night_routine"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

// AnalyzeSkin analyzes user's skin condition and generates recommendations
func AnalyzeSkin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := middleware.GetUserID(c)

	var req SkinAnalysisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate recommendations
	morningRoutine := generateMorningRoutine(req.SkinType, req.SkinIssues, req.Sensitivity)
	nightRoutine := generateNightRoutine(req.SkinType, req.SkinIssues, req.Sensitivity)

	// Create analysis
	analysis := models.SkinAnalysis{
		UserID:         userID,
		SkinType:       req.SkinType,
		SkinIssues:     req.SkinIssues,
		Sensitivity:    req.Sensitivity,
		MorningRoutine: morningRoutine,
		NightRoutine:   nightRoutine,
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
	}

	if err := models.CreateAnalysis(db, &analysis); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create analysis"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Skin analysis created successfully",
		"data": SkinAnalysisResponse{
			ID:             analysis.ID,
			SkinType:       analysis.SkinType,
			SkinIssues:     req.SkinIssues,
			Sensitivity:    analysis.Sensitivity,
			MorningRoutine: analysis.MorningRoutine,
			NightRoutine:   analysis.NightRoutine,
			CreatedAt:      analysis.CreatedAt,
			UpdatedAt:      analysis.UpdatedAt,
		},
	})
}

// GetHistory gets user's skin analysis history
func GetHistory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := middleware.GetUserID(c)

	analyses, err := models.GetAnalysisByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get history"})
		return
	}

	var response []SkinAnalysisResponse
	for _, analysis := range analyses {
		response = append(response, SkinAnalysisResponse{
			ID:             analysis.ID,
			SkinType:       analysis.SkinType,
			SkinIssues:     analysis.SkinIssues,
			Sensitivity:    analysis.Sensitivity,
			MorningRoutine: analysis.MorningRoutine,
			NightRoutine:   analysis.NightRoutine,
			CreatedAt:      analysis.CreatedAt,
			UpdatedAt:      analysis.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// GetAnalysisDetail gets a specific analysis
func GetAnalysisDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := middleware.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	analysis, err := models.GetAnalysisByID(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Analysis not found"})
		return
	}

	// Check if analysis belongs to user
	if analysis.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": SkinAnalysisResponse{
			ID:             analysis.ID,
			SkinType:       analysis.SkinType,
			SkinIssues:     analysis.SkinIssues,
			Sensitivity:    analysis.Sensitivity,
			MorningRoutine: analysis.MorningRoutine,
			NightRoutine:   analysis.NightRoutine,
			CreatedAt:      analysis.CreatedAt,
			UpdatedAt:      analysis.UpdatedAt,
		},
	})
}

// Helper functions for generating routines

func generateMorningRoutine(skinType string, issues []string, sensitivity string) models.RoutineData {
	routine := models.RoutineData{
		Steps: []string{},
		Tips:  "",
	}

	// Base morning routine
	routine.Steps = append(routine.Steps, "1. Pembersih wajah - Cuci wajah dengan pembersih yang sesuai dengan jenis kulit")

	// Add specific steps based on skin type
	switch skinType {
	case "kering":
		routine.Steps = append(routine.Steps, "2. Toner - Gunakan toner untuk melembabkan")
		routine.Steps = append(routine.Steps, "3. Essence - Aplikasikan essence hydrating")
	case "berminyak":
		routine.Steps = append(routine.Steps, "2. Toner - Gunakan toner oil-control")
		routine.Steps = append(routine.Steps, "3. Clay mask ringan - Aplikasikan 1-2x seminggu")
	case "kombinasi":
		routine.Steps = append(routine.Steps, "2. Toner - Gunakan toner balancing")
		routine.Steps = append(routine.Steps, "3. Serum - Aplikasikan serum yang ringan")
	case "sensitif":
		routine.Steps = append(routine.Steps, "2. Toner - Gunakan toner yang gentle dan bebas alkohol")
		routine.Steps = append(routine.Steps, "3. Essence - Aplikasikan essence yang hypoallergenic")
	}

	// Handle specific issues
	hasAcne := contains(issues, "jerawat")
	hasDull := contains(issues, "kusam")
	hasSpot := contains(issues, "flek hitam")

	if hasAcne {
		routine.Steps = append(routine.Steps, "4. Acne treatment - Aplikasikan produk mengandung salicylic acid atau benzoyl peroxide")
	}

	if hasDull {
		routine.Steps = append(routine.Steps, "4. Vitamin C serum - Untuk mencerahkan wajah")
	}

	if hasSpot {
		routine.Steps = append(routine.Steps, "4. Spot treatment - Gunakan produk yang mengandung niacinamide atau vitamin C")
	}

	// Moisturizer step
	if skinType == "kering" {
		routine.Steps = append(routine.Steps, "5. Moisturizer - Gunakan heavy moisturizer")
	} else if skinType == "berminyak" {
		routine.Steps = append(routine.Steps, "5. Moisturizer - Gunakan lightweight moisturizer atau gel")
	} else {
		routine.Steps = append(routine.Steps, "5. Moisturizer - Gunakan moisturizer yang sesuai")
	}

	// Sunscreen
	routine.Steps = append(routine.Steps, "6. Sunscreen - Aplikasikan sunscreen SPF 30+ sebagai langkah terakhir")

	// Tips
	if sensitivity == "tinggi" {
		routine.Tips = "Gunakan produk hypoallergenic dan lakukan patch test sebelum pemakaian. Hindari produk yang mengandung fragrance atau alcohol"
	} else if sensitivity == "sedang" {
		routine.Tips = "Hindari penggunaan produk yang terlalu aktif di pagi hari. Fokus pada pembersihan dan perlindungan"
	} else {
		routine.Tips = "Pagi adalah waktu terbaik untuk menggunakan vitamin C dan sunscreen untuk perlindungan maksimal"
	}

	return routine
}

func generateNightRoutine(skinType string, issues []string, sensitivity string) models.RoutineData {
	routine := models.RoutineData{
		Steps: []string{},
		Tips:  "",
	}

	// Base night routine
	routine.Steps = append(routine.Steps, "1. Makeup remover - Gunakan makeup remover yang efektif")
	routine.Steps = append(routine.Steps, "2. Pembersih wajah - Cuci wajah dengan pembersih kedua (double cleanse)")

	// Add specific steps based on skin type
	switch skinType {
	case "kering":
		routine.Steps = append(routine.Steps, "3. Toner - Gunakan toner yang melembabkan")
		routine.Steps = append(routine.Steps, "4. Essence - Aplikasikan essence yang nourishing")
	case "berminyak":
		routine.Steps = append(routine.Steps, "3. Toner - Gunakan toner oil-control")
	case "kombinasi":
		routine.Steps = append(routine.Steps, "3. Toner - Gunakan toner balancing")
	case "sensitif":
		routine.Steps = append(routine.Steps, "3. Toner - Gunakan toner yang gentle")
	}

	// Handle specific issues - night is better for treatments
	hasAcne := contains(issues, "jerawat")
	hasDull := contains(issues, "kusam")
	hasSpot := contains(issues, "flek hitam")
	hasWrinkles := contains(issues, "kerutan")

	if hasAcne {
		routine.Steps = append(routine.Steps, "4. Acne treatment - Aplikasikan produk dengan retinoid atau benzoyl peroxide")
	}

	if hasDull || hasWrinkles {
		routine.Steps = append(routine.Steps, "4. Retinol/Retinoid - Gunakan untuk anti-aging dan brightening")
	}

	if hasSpot {
		routine.Steps = append(routine.Steps, "4. Niacinamide serum - Untuk mengurangi hyperpigmentation")
	}

	// Add serums and treatments
	routine.Steps = append(routine.Steps, "5. Serum/Treatment - Aplikasikan serum sesuai kebutuhan")

	// Night moisturizer
	if skinType == "kering" {
		routine.Steps = append(routine.Steps, "6. Night cream - Gunakan night cream yang kaya dan nourishing")
	} else if skinType == "berminyak" {
		routine.Steps = append(routine.Steps, "6. Night moisturizer - Gunakan moisturizer yang ringan")
	} else {
		routine.Steps = append(routine.Steps, "6. Night moisturizer - Gunakan moisturizer yang sesuai")
	}

	// Optional: sleeping mask
	if skinType == "kering" {
		routine.Steps = append(routine.Steps, "7. Sleeping mask - Opsional untuk intensif hydration")
	}

	// Tips
	if sensitivity == "tinggi" {
		routine.Tips = "Malam adalah waktu untuk treatment intensif. Hindari produk yang terlalu aktif jika kulit terasa sensitif"
	} else if sensitivity == "sedang" {
		routine.Tips = "Gunakan waktu malam untuk ingredient aktif seperti retinoid dan acid. Mulai dengan konsentrasi rendah"
	} else {
		routine.Tips = "Malam adalah waktu ideal untuk menggunakan bahan aktif seperti retinol, vitamin A, dan exfoliating acids"
	}

	return routine
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
