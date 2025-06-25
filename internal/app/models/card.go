package models

import (
	"gorm.io/gorm"
	"time"
)

type CardDetails struct {
	gorm.Model

	ExpireDate time.Time `json:"expire_date" gorm:"type:date"`
	IssueDate  time.Time `json:"issue_date" gorm:"type:date"`
	CardType   string    `json:"card_type" gorm:"type:text"`
	Code       string    `json:"-" gorm:"type:text"`
	InBalance  float64   `json:"in_balance" gorm:"type:numeric"`
	DebtOsd    float64   `json:"debt_osd" gorm:"type:numeric"`
	DebtOsk    float64   `json:"debt_osk" gorm:"type:numeric"`
	OutBalance float64   `json:"out_balance" gorm:"type:numeric"`
	Coast      float64   `json:"coast" gorm:"type:numeric"`

	WorkerID uint   `json:"worker_id" gorm:"not null"`
	Worker   Worker `json:"worker" gorm:"foreignkey:WorkerID"`
}
