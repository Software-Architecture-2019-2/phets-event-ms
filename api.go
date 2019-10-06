package main

import (
	"github.com/Software-Architecture-2019-2/sa-event-ms/resource"

	"github.com/gorilla/mux"
)

// Router with REST API routes for Event Microservice
var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/event", resource.GetAllEventsEndpoint).Methods("GET")
	Router.HandleFunc("/event/{id}", resource.GetEventEndpoint).Methods("GET")
	Router.HandleFunc("/event", resource.CreateEventEndpoint).Methods("POST")
	Router.HandleFunc("/event/{id}", resource.DeleteEventEndpoint).Methods("DELETE")
	Router.HandleFunc("/event/{id}", resource.UpdateEventEndpoint).Methods("PUT")
}
