package main

import (
	"log"
	"net/http"

	"github.com/boytur/go-crud-mvc/config"
	"github.com/boytur/go-crud-mvc/controllers"
	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
