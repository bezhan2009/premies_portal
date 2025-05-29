package models

import "gorm.io/gorm"

type Office struct {
	gorm.Model

	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Director    *User  `gorm:"foreignKey:DirectorID"` // Связь с директором
	DirectorID  *int   // внешний ключ на User
}

type OfficeUser struct {
	gorm.Model

	OfficeID int
	Office   Office `gorm:"foreignkey:OfficeID"`
	WorkerID int
	Worker   Worker `gorm:"foreignkey:WorkerID"`
}
