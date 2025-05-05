package models

import "gorm.io/gorm"

type CardTurnovers struct {
	gorm.Model

	ActiveCards float64
	Month       float64

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
