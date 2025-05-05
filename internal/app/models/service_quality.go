package models

import "gorm.io/gorm"

type ServiceQuality struct {
	gorm.Model

	AverageScore float64
	Coefficient  float64

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
