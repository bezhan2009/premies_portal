package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(255);unique;not null"`
	Email    string `gorm:"type:varchar(255);unique;"`
	Phone    string `gorm:"type:varchar(9);unique;"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`

	FullName string `json:"full_name" gorm:"type:varchar(255);"`

	RoleID int  `json:"role_id" gorm:"not null"`
	Role   Role `json:"-" gorm:"foreignKey:RoleID"`

	CardTurnovers   []CardTurnovers   `gorm:"foreignKey:WorkerID"`
	CardSales       []CardSales       `gorm:"foreignKey:WorkerID"`
	OperatingActive []OperatingActive `gorm:"foreignKey:WorkerID"`
	ServiceQuality  []ServiceQuality  `gorm:"foreignKey:WorkerID"`
}

type Role struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(20);unique;not null"`
}
