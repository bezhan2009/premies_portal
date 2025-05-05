package models

import "gorm.io/gorm"

type CardSales struct {
	gorm.Model

	Cards         float64
	MobileBank    float64
	Overdraft     float64
	SalaryProject float64

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
