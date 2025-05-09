package models

import "gorm.io/gorm"

type OperatingActive struct {
	gorm.Model

	Deposits  float64 `json:"deposits" gorm:"default:0.0"`
	PiggyBank float64 `json:"piggy_bank" gorm:"default:0.0"`
	Transfers float64 `json:"transfers" gorm:"default:0.0"`

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
