package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *director `json:"director"`

}

type director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []movie

func getmovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies {

		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range movies {

		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createmovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main(){
	r := mux.NewRouter()

    movies = append(movies, movie{ID: "1",Isbn: "438227",Title: "movie one",Director: &director{Firstname: "john",Lastname: "doe"}})
	movies = append(movies, movie{ID: "2",Isbn: "43866",Title: "movie two",Director: &director{Firstname: "steve",Lastname: "smith"}})
	r.HandleFunc("/movies",getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getmovie).Methods("GET")
	r.HandleFunc("/movies",createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deletemovie).Methods("DELETE")
	

	fmt.Printf("start server at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))

}