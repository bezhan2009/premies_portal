package models

import "gorm.io/gorm"

type ServiceQuality struct {
	gorm.Model

	AverageScore float64 `json:"average_score" gorm:"default:0.0"`
	Coefficient  float64 `json:"coefficient" gorm:"default:0.0"`
	Complaint    float64 `json:"complaint" gorm:"default:0.0"`
	Tests        float64 `json:"tests" gorm:"default:0.0"`
	Bonus        float64 `json:"bonus" gorm:"default:0.0"`

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
