package models

import "gorm.io/gorm"

type CardSales struct {
	gorm.Model

	CardsPrem     float64 `json:"cards_prem" gorm:"default:0.0"`
	SalaryProject float64 `json:"salary_project" gorm:"default:0.0"`

	WorkerID uint   `gorm:"not null"`
	Worker   Worker `json:"-" gorm:"foreignkey:WorkerID"`
}
