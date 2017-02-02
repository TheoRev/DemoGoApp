package models

import "github.com/jinzhu/gorm"

// User estructura para la autentificacion de usuarios
type User struct {
	gorm.Model
	User     string `json:"user" gorm:"not null;unique"`
	Password string `json:"password,omitempty" gorm:"not null"`
}
