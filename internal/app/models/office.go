package models

import "gorm.io/gorm"

type Office struct {
	gorm.Model

	Name       string `gorm:"type:varchar(255)"`
	Director   *User  `gorm:"foreignKey:DirectorID"` // Связь с директором
	DirectorID *int   // внешний ключ на User
}

type OfficeUser struct {
	gorm.Model

	OfficeID int
	Office   Office `gorm:"foreignKey:OfficeID"`
	UserID   int
	User     User `gorm:"foreignKey:UserID"`
}
