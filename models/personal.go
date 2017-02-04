package models

// Personal structura con los datos principales del personal
type Personal struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Dni         string `json:"dni"`
	Address     string `json:"address"`
}
