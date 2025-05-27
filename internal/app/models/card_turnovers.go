package models

import "gorm.io/gorm"

type CardTurnovers struct {
	gorm.Model

	CardTurnoversPrem float64 `json:"card_turnovers_prem"`
	ActiveCardsPerms  float64 `json:"active_cards_perms"`

	WorkerID uint   `gorm:"not null"`
	Worker   Worker `json:"-" gorm:"foreignkey:WorkerID"`
}
