package dao

import (
	"DemoGoApp/util"
	"database/sql"
	"fmt"
	"log"
	// Importa el driver de postgresql
	_ "github.com/lib/pq"
)

// Get obtiene la conexion a postgres
func Get() *sql.DB {
	config, err := util.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
