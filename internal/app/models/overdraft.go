package models

import "gorm.io/gorm"

type Overdraft struct {
	gorm.Model

	OverdraftPerm float64 `json:"overdraft_perm" gorm:"default:0.0"`
	WorkerID      uint    `json:"worker_id"`
	Worker        Worker  `gorm:"foreignKey: WorkerID" json:"worker"`
}
