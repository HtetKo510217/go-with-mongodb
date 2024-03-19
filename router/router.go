package router

import (
	"github.com/gorilla/mux"
	"github.com/htetko/go-with-mongodb/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}