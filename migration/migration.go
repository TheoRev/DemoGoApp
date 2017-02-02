package migration

import (
	"DemoGoApp/configuration"
	"DemoGoApp/models"
)

// Migrate de la estructura de DB
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Personal{})
}
