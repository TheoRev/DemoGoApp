package models

// User estructura para la autentificacion de usuarios
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
