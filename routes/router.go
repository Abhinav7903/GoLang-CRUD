package routes

import (
	"net/http"

	"github.com/Abhinav7903/mongo/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/api/GetMyAllMovies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controller.MarksAsWatched).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")
	
	http.ListenAndServe(":8080", r)
	return r
}	