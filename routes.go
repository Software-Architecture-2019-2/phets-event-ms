package main

import (
	"github.com/Software-Architecture-2019-2/phets-event-ms/resource"

	"github.com/gorilla/mux"
)

// Router with REST API routes for Event Microservice
var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/event", resource.GetAllEvents).Methods("GET")
	Router.HandleFunc("/event/{id}", resource.GetEventByID).Methods("GET")
	Router.HandleFunc("/event", resource.CreateEvent).Methods("POST")
	Router.HandleFunc("/event/{id}", resource.DeleteEvent).Methods("DELETE")
	Router.HandleFunc("/event/{id}", resource.UpdateEvent).Methods("PUT")
}
