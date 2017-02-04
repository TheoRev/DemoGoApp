package interfaces

import "DemoGoApp/models"

// UserDAO interface con los metodos del CRUD
type UserDAO interface {
	Create(u *models.User) error
	// Update(u *models.User) error
	// Delete(id int) error
	FindBy(u *models.User) ([]models.User, error)
	// FindById(id int) (models.User, error)
	// FindAll() ([]models.User, error)
}
