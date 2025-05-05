package models

import "gorm.io/gorm"

type OperatingActive struct {
	gorm.Model

	Deposits  float64
	PiggyBank float64
	Transfers float64

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
