package models

// Configuration structura con los parametros de conexion a la DB
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}
