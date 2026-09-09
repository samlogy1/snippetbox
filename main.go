package main

import (
	"fmt"
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the other side"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet money..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet money..."))
}

func main() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :9000")


	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)

	fmt.Println("Hello World!")
}