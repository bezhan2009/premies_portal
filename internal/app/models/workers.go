package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model

	Salary        float32 `gorm:"not null"`
	Position      string  `json:"position"`
	Plan          uint    `json:"plan" gorm:"default:0"`
	SalaryProject uint    `json:"salary_project" gorm:"default:0"`

	UserID uint `json:"user_id" gorm:"default:0"`
	User   User `json:"-"`

	CardTurnovers  []CardTurnovers  `gorm:"foreignKey:WorkerID"`
	CardSales      []CardSales      `gorm:"foreignKey:WorkerID"`
	ServiceQuality []ServiceQuality `gorm:"foreignKey:WorkerID"`
}
