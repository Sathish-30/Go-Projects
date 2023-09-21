package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func main() {

	PORT_NUMBER := ":3000"
	router := mux.NewRouter()

	// Home router (Health check)
	router.HandleFunc("/",func(w http.ResponseWriter , r *http.Request){
		fmt.Fprint(w,"Hello world")
	})

	// Get all movies
	router.HandleFunc("/movies",handleGetAllMovies).Methods("GET")
	
	// Get movie by Id
	router.HandleFunc("/movies/{id}" , handleGetMovie).Methods("GET")

	// Post / Create Movie
	router.HandleFunc("/movies",handleCreateMovie).Methods("POST")

	// Update a movie via ID
	router.HandleFunc("/movies/{id}" , handleUpdateMovies).Methods("PUT")

	// Delete movie via ID
	router.HandleFunc("/movies/{id}" , handleDeleteMovie).Methods("DELETE")

	fmt.Printf("Server is starting at the Port %v " , PORT_NUMBER)

	if err := http.ListenAndServe(PORT_NUMBER, nil) ; err != nil{
		log.Fatal(err)
	}
}