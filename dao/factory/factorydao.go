package factory

import (
	"DemoGoApp/dao/interfaces"
	"DemoGoApp/dao/psql"
	"log"
)

func FactoryDAO(e string) interfaces.UserDAO {
	var i interfaces.UserDAO
	switch e {
	case "postgres":
		i = psql.UserImpl{}
	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}
	return i
}
