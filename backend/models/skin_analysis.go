package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type SkinAnalysis struct {
	ID             uint        `json:"id" gorm:"primaryKey"`
	UserID         uint        `json:"user_id" gorm:"not null;index"`
	User           User        `json:"-" gorm:"foreignKey:UserID"`
	SkinType       string      `json:"skin_type" gorm:"not null"` // kering, berminyak, kombinasi, sensitif
	SkinIssuesJSON string      `json:"-" gorm:"type:json;column:skin_issues"`
	Sensitivity    string      `json:"sensitivity" gorm:"not null"` // rendah, sedang, tinggi
	MorningRoutineJSON string   `json:"-" gorm:"type:json;column:morning_routine"`
	NightRoutineJSON   string   `json:"-" gorm:"type:json;column:night_routine"`
	CreatedAt      int64       `json:"created_at"`
	UpdatedAt      int64       `json:"updated_at"`
	// Virtual fields untuk JSON response
	SkinIssues     []string    `json:"skin_issues" gorm:"-"`
	MorningRoutine RoutineData `json:"morning_routine" gorm:"-"`
	NightRoutine   RoutineData `json:"night_routine" gorm:"-"`
}

type RoutineData struct {
	Steps []string `json:"steps"`
	Tips  string   `json:"tips"`
}

func (SkinAnalysis) TableName() string {
	return "skin_analysis"
}

// AfterFind hook untuk unmarshal JSON
func (sa *SkinAnalysis) AfterFind(tx *gorm.DB) error {
	if sa.SkinIssuesJSON != "" {
		json.Unmarshal([]byte(sa.SkinIssuesJSON), &sa.SkinIssues)
	}
	if sa.MorningRoutineJSON != "" {
		json.Unmarshal([]byte(sa.MorningRoutineJSON), &sa.MorningRoutine)
	}
	if sa.NightRoutineJSON != "" {
		json.Unmarshal([]byte(sa.NightRoutineJSON), &sa.NightRoutine)
	}
	return nil
}

// BeforeSave hook untuk marshal JSON
func (sa *SkinAnalysis) BeforeSave(tx *gorm.DB) error {
	if len(sa.SkinIssues) > 0 {
		data, _ := json.Marshal(sa.SkinIssues)
		sa.SkinIssuesJSON = string(data)
	}
	if len(sa.MorningRoutine.Steps) > 0 {
		data, _ := json.Marshal(sa.MorningRoutine)
		sa.MorningRoutineJSON = string(data)
	}
	if len(sa.NightRoutine.Steps) > 0 {
		data, _ := json.Marshal(sa.NightRoutine)
		sa.NightRoutineJSON = string(data)
	}
	return nil
}

// Get analysis by ID
func GetAnalysisByID(db *gorm.DB, id uint) (*SkinAnalysis, error) {
	var analysis SkinAnalysis
	if err := db.Where("id = ?", id).First(&analysis).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

// Get all analysis by user ID
func GetAnalysisByUserID(db *gorm.DB, userID uint) ([]SkinAnalysis, error) {
	var analyses []SkinAnalysis
	if err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&analyses).Error; err != nil {
		return nil, err
	}
	return analyses, nil
}

// Create analysis
func CreateAnalysis(db *gorm.DB, analysis *SkinAnalysis) error {
	return db.Create(analysis).Error
}
