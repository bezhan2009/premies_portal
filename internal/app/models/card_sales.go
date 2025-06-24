package models

import "gorm.io/gorm"

type CardSales struct {
	gorm.Model

	DebOsd        float64 `json:"deb_osd"`
	CardsSailed   uint    `json:"cards_sailed"`
	CardsPrem     float64 `json:"cards_prem" gorm:"default:0.0"`
	SalaryProject float64 `json:"salary_project" gorm:"default:0.0"`

	WorkerID uint   `gorm:"not null"`
	Worker   Worker `json:"-" gorm:"foreignkey:WorkerID"`
}
