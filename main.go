package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")
	
	w.Write([]byte("Hello from the other side\n"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}

	msg := fmt.Sprintf("Display a specific snippet with with ID %d.....\n", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet money...\n"))
}
// func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// category := r.PathValue("category")
	// itemID := r.PathValue("itemID")
// }

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet....\n"))
}

func main() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on http://localhost:9000")


	err := http.ListenAndServe(":9000", mux)
	log.Print("This server was built by Akeju Samuel")
	log.Fatal(err)

	fmt.Println("Hello World!")
}