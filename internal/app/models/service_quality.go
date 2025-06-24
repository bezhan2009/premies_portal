package models

import "gorm.io/gorm"

type ServiceQuality struct {
	gorm.Model

	CallCenter  float64 `json:"call_center" gorm:"default:0.0"`
	Coefficient float64 `json:"coefficient" gorm:"default:0.0"`
	Complaint   float64 `json:"complaint" gorm:"default:0.0"`
	Tests       float64 `json:"tests" gorm:"default:0.0"`

	WorkerID uint   `gorm:"not null"`
	Worker   Worker `json:"-" gorm:"foreignkey:WorkerID"`
}
