package models

import "gorm.io/gorm"

type CardTurnovers struct {
	gorm.Model

	CardTurnoversPrem float64 `json:"active_cards_prem" gorm:"not null"`

	WorkerID uint `gorm:"not null"`
	Worker   User `json:"-" gorm:"foreignkey:WorkerID"`
}
