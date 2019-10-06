package main

import (
	"github.com/Software-Architecture-2019-2/sa-event-ms/resource"

	"github.com/gorilla/mux"
)

// Router with REST API routes for Event Microservice
var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/people", resource.GetAllEventsEndpoint).Methods("GET")
	Router.HandleFunc("/people/{id}", resource.GetEventEndpoint).Methods("GET")
	Router.HandleFunc("/people", resource.CreateEventEndpoint).Methods("POST")
	Router.HandleFunc("/people/{id}", resource.DeleteEventEndpoint).Methods("DELETE")
	Router.HandleFunc("/people/{id}", resource.UpdateEventEndpoint).Methods("PUT")
}
