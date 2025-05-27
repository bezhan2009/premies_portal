package models

import "gorm.io/gorm"

type History struct {
	gorm.Model

	FinalAward float64

	CardSalesID       int
	CardTurnoversID   int
	OperatingActiveID int
	ServiceQualityID  int

	CardSales      CardSales      `gorm:"foreignKey:CardSalesID"`
	CardTurnovers  CardTurnovers  `gorm:"foreignKey:CardTurnoversID"`
	ServiceQuality ServiceQuality `gorm:"foreignKey:ServiceQualityID"`

	WorkerID uint `gorm:"not null"`
	Worker   User `gorm:"foreignkey:WorkerID"`
}
