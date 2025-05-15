package models

import "gorm.io/gorm"

type Overdraft struct {
	gorm.Model

	OverdraftCount uint `json:"overdraft_count"`
}
