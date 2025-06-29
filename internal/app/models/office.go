package models

import "gorm.io/gorm"

type Office struct {
	gorm.Model

	Title       string `json:"title" gorm:"type:varchar(255);unique"`
	Description string `json:"description" gorm:"type:text"`
	DirectorID  *int   `json:"director_id"`
	Director    *User  `json:"-" gorm:"foreignKey:DirectorID"`

	OfficeUsers []OfficeUser `json:"office_user" gorm:"foreignkey:OfficeID"`
}

type OfficeUser struct {
	gorm.Model

	OfficeID int    `json:"office_id"`
	Office   Office `json:"-" gorm:"foreignkey:OfficeID"`
	WorkerID int    `json:"worker_id"`
	Worker   Worker `json:"worker" gorm:"foreignkey:WorkerID"`
}
