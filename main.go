package main

import (
	"FirstUseCase/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Handle("/GetArtistByCountry/{countryName}", handlers.GetArtistByCountry()).Methods("GET")
	router.Handle("/GetSimilarTracks/{artistName}/{trackName}", handlers.GetSimilarTracks()).Methods("GET")
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Listening on 8080")
	server.ListenAndServe()
}
