package main

import (
	"DemoGoApp/dao/factory"
	"DemoGoApp/models"
	"DemoGoApp/util"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	initCrud()
}

func initCrud() {
	config, err := util.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	userDAO := factory.FactoryDAO(config.Engine)
	user := models.User{}

	fmt.Print("Ingrese su UserName: ")
	fmt.Scan(&user.UserName)
	fmt.Print("Ingrese su Password: ")
	fmt.Scan(&user.Password)

	err = userDAO.Create(&user)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
}

func initServer() {
	// route := mux.NewRouter().StrictSlash(false)
	// fs := http.FileServer(http.Dir("./public"))
	// route.Handle("/", fs)

	route := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	route.Handle("/", fs)

	server := &http.Server{
		Addr:           ":1700",
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening http://localhost:1700")
	log.Fatal(server.ListenAndServe())
}
