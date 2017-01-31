package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public"))
	r.Handle("/", fs)
	r.Handle("/css/", fs)
	r.Handle("/js/", fs)

	server := &http.Server{
		Addr:           ":1700",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening http://localhost:1700")
	log.Fatal(server.ListenAndServe())
}
