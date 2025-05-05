package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(255);unique;not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Phone    string `gorm:"type:varchar(9);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`

	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`

	RoleID int
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
