package models

import "github.com/jinzhu/gorm"

// Personal structura con los datos principales del personal
type Personal struct {
	gorm.Model
	FirstName   string `json:"firstName" gorm:"not null"`
	LastName    string `json:"lastName" gorm:"not null"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null;type:varchar(9)"`
	Dni         string `json:"dni" gorm:"not null;type:varchar(8)"`
	Adress      string `json:"adress" gorm:"type:varchar(150)"`
}
