package psql

import (
	"DemoGoApp/dao"
	"DemoGoApp/models"
)

// UserImpl implementacioon de la interace
type UserImpl struct{}

// Create para ingresar nuevos usuarios
func (daou UserImpl) Create(u *models.User) error {
	query := "INSERT INTO users (user_name, password) VALUES($1, $2)"
	db := dao.Get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.UserName, u.Password)
	row.Scan(&u.ID)
	return nil
}

// FindBy para autentificar usuarios
func (daou UserImpl) FindBy(u *models.User) ([]models.User, error) {
	query := "SELECT user_name, password FROM users where user_name=$1 and password=$2"
	users := make([]models.User, 0)
	db := dao.Get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.UserName, &row.Password)
		if err != nil {
			return users, err
		}
		users = append(users, row)
	}
	return users, nil
}
