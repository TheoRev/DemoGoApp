package main

import (
	"DemoGoApp/migration"
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	migrar()
}

func migrar() {
	log.Println("Entró al metodo migrar()")
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Inició la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración...")
	}
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
