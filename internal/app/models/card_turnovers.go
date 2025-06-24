package models

import "gorm.io/gorm"

type CardTurnovers struct {
	gorm.Model

	ActivatedCards    uint    `json:"activated_cards"`
	CardTurnoversPrem float64 `json:"card_turnovers_prem"`
	ActiveCardsPerms  float64 `json:"active_cards_perms"`

	WorkerID uint   `gorm:"not null"`
	Worker   Worker `json:"-" gorm:"foreignkey:WorkerID"`
}
