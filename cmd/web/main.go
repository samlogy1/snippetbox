package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetview)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)



	log.Println("starting a new server on port: http://localhost:9000")
	log.Println("this server was built by Akeju Samuel")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}