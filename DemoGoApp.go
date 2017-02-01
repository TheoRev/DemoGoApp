package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
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
