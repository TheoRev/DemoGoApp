package conexion

import (
	"encoding/json"
	"log"
	"os"
)

// Conexion estructura con los parametros de conexion a la DB
type Conexion struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConexion obtiene la conexion a la base de datos
func GetConexion() Conexion {
	var conexion Conexion
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&conexion)
	if err != nil {
		log.Fatal(err)
	}

	return conexion
}
