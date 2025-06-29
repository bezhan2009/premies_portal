package models

import "gorm.io/gorm"

type MobileBankDetails struct {
	gorm.Model

	INN  string  `json:"inn" gorm:"varchar(20)"`
	Prem float64 `json:"prem"`

	WorkerID uint   `gorm:"not null"`
	Worker   Worker `json:"-" gorm:"foreignkey:WorkerID"`
}
