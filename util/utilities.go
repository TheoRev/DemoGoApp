package util

import (
	"DaoDesign/models"
	"encoding/json"
	"log"
	"os"
)

// GetConfiguration obtiene los parametros de conexion
func GetConfiguration() (models.Configuration, error) {
	config := models.Configuration{}
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
		return config, err
	}

	return config, nil
}
