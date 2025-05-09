package models

import "gorm.io/gorm"

type ServiceQuality struct {
	gorm.Model

	AverageScore float64 `json:"average_score" gorm:"default:0.0"`
	Coefficient  float64 `json:"coefficient" gorm:"default:0.0"`

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
