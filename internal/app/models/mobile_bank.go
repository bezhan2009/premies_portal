package models

import "gorm.io/gorm"

type MobileBankSales struct {
	gorm.Model

	MobileBankConnects uint    `json:"mobile_bank_connects"`
	MobileBankPrem     float64 `json:"mobile_bank_prem" gorm:"default:0.0"`
	WorkerID           uint    `json:"worker_id" gorm:"not null"`
	Worker             Worker  `json:"-" gorm:"foreignkey:WorkerID"`
}
