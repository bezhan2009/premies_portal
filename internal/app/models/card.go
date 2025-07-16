package models

import (
	"time"

	"gorm.io/gorm"
)

type CardDetails struct {
	gorm.Model

	ExpireDate   time.Time `json:"expire_date" gorm:"type:date"`
	IssueDate    time.Time `json:"issue_date" gorm:"type:date"`
	CardType     string    `json:"card_type" gorm:"type:text"`
	Code         string    `json:"code" gorm:"type:text"`
	InBalance    float64   `json:"in_balance" gorm:"type:numeric"`
	DebtOsd      float64   `json:"debt_osd" gorm:"type:numeric"`
	DebtOsk      float64   `json:"debt_osk" gorm:"type:numeric"`
	OutBalance   float64   `json:"out_balance" gorm:"type:numeric"`
	CoastCards   float64   `json:"coast" gorm:"type:numeric"`
	CoastCredits float64   `json:"coast_credits" gorm:"type:numeric"`

	WorkerID uint   `json:"worker_id" gorm:"not null"`
	Worker   Worker `json:"worker" gorm:"foreignkey:WorkerID"`

	OwnerName string `json:"owner_name"`
}

type CardsCharters struct {
	DebtOsd        float64 `json:"debt_osd"`
	DebtOsk        float64 `json:"debt_osk"`
	OutBalance     float64 `json:"out_balance"`
	InBalance      float64 `json:"in_balance"`
	CardsInGeneral uint    `json:"cards_in_general"`
	CardsForMonth  uint    `json:"cards_for_month"`
	ActivatedCards uint    `json:"activated_cards"`
}
