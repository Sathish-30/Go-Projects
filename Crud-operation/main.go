package main

import (
	"encoding/json"
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

	addMovies()
	// Home router (Health check)
	router.HandleFunc("/",handleHomeRoute)

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

// Add function add movies to the Movie slice
func addMovies(){
	movies = append(movies, Movie{ID:"1" , Isbn: "4387" , Title: "Leo" , Director: &Director{FirstName: "Lokesh" , LastName: "Kanagaraj"} })
	movies = append(movies, Movie{ID:"2" , Isbn: "5367" , Title: "Theri" , Director: &Director{FirstName: "Atlee" , LastName: "Guna"} })
	movies = append(movies, Movie{ID:"3" , Isbn: "8374" , Title: "Baasha" , Director: &Director{FirstName: "rajini" , LastName: "kanth"} })
}

// This request get trigger when the route is in / or home route
func handleHomeRoute(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w,"Hello world")
}

func handleGetAllMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(movies)
}

func handleDeleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for index , movie := range movies{
		if movie.ID == id {
			// Where the ... operator will convert the slice into a many single element
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}