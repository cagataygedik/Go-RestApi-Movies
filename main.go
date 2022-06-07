package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Imdb     string    `json:"imdb"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	//delete the movie with the id
	//add a new movie = the movie we sent in the body of postman
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Imdb: "9.2", Title: "The Godfather ", Director: &Director{Firstname: "Francis Ford", Lastname: "Coppola"}})
	movies = append(movies, Movie{ID: "2", Imdb: "8.9", Title: "12 Angry Men", Director: &Director{Firstname: "Sidney", Lastname: "Lumet"}})
	movies = append(movies, Movie{ID: "3", Imdb: "8.5", Title: "Casablanca", Director: &Director{Firstname: "Michael", Lastname: "Curtiz"}})
	movies = append(movies, Movie{ID: "4", Imdb: "8.4", Title: "Rear Window", Director: &Director{Firstname: "Alfred", Lastname: "Hitchcock"}})
	movies = append(movies, Movie{ID: "5", Imdb: "8.5", Title: "City Lights", Director: &Director{Firstname: "Charles", Lastname: "Chaplin"}})
	movies = append(movies, Movie{ID: "6", Imdb: "8.9", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})
	movies = append(movies, Movie{ID: "7", Imdb: "8.9", Title: "The Lord of the Rings: The Return of the King", Director: &Director{Firstname: "Peter", Lastname: "Jackson"}})
	movies = append(movies, Movie{ID: "8", Imdb: "8.9", Title: "Schindler's List", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: "9", Imdb: "8.3", Title: "Citizen Kane", Director: &Director{Firstname: "Orson", Lastname: "Welles"}})
	movies = append(movies, Movie{ID: "10", Imdb: "8.3", Title: "Vertigo", Director: &Director{Firstname: "Alfred", Lastname: "Hitchcock"}})
	movies = append(movies, Movie{ID: "11", Imdb: "8.5", Title: "Psycho", Director: &Director{Firstname: "Alfred", Lastname: "Hitchcock"}})
	movies = append(movies, Movie{ID: "12", Imdb: "8.3", Title: "Singin' in the Rain", Director: &Director{Firstname: "Gene", Lastname: "Kelly"}})
	movies = append(movies, Movie{ID: "13", Imdb: "8.4", Title: "Dr. Strangelove or: How I Learned to Stop Worrying and Love the Bomb", Director: &Director{Firstname: "Stanley", Lastname: "Kubrick"}})
	movies = append(movies, Movie{ID: "14", Imdb: "8.3", Title: "North by Northwest", Director: &Director{Firstname: "Alfred", Lastname: "Hitchcock"}})
	movies = append(movies, Movie{ID: "15", Imdb: "8.5", Title: "Modern Times", Director: &Director{Firstname: "Charles", Lastname: "Chaplin"}})
	movies = append(movies, Movie{ID: "16", Imdb: "8.1", Title: "Sweet Smell of Success", Director: &Director{Firstname: "Alexander", Lastname: "Mackendrick"}})
	movies = append(movies, Movie{ID: "17", Imdb: "8.8", Title: "The Lord of the Rings: The Fellowship of the Ring", Director: &Director{Firstname: "Çağrı", Lastname: "Gedik"}})
	movies = append(movies, Movie{ID: "18", Imdb: "9.0", Title: "The Godfather: Part II", Director: &Director{Firstname: "Francis Ford", Lastname: "Coppola"}})
	movies = append(movies, Movie{ID: "19", Imdb: "8.0", Title: "The Wizard of Oz", Director: &Director{Firstname: "Victor", Lastname: "Fleming"}})
	movies = append(movies, Movie{ID: "20", Imdb: "8.2", Title: "All About Eve ", Director: &Director{Firstname: "Joseph L. ", Lastname: "Mankiewicz"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
