package main

import (
	"./handlers"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/Build", handlers.Build)

	log.Println("Server running in port number 1094")
	err := http.ListenAndServe(":1094", mux)

	if err != nil {
		log.Fatalf("Something happend", "%s", err)
	}

}
