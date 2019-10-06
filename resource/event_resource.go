package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Software-Architecture-2019-2/sa-event-ms/model"
	"github.com/Software-Architecture-2019-2/sa-event-ms/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var people []model.Event

// GetEventEndpoint gets an event
func GetEventEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	params := mux.Vars(r)
	eventID, _ := primitive.ObjectIDFromHex(params["id"])
	event := service.GetEvent(eventID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

// GetAllEventsEndpoint gets all events
func GetAllEventsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	events := service.GetAllEvents()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

// CreateEventEndpoint creates an event
func CreateEventEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	var event model.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to create")
	}
	json.Unmarshal(reqBody, &event)

	insertedEvent := service.InsertOneValue(event)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedEvent)
}

// DeleteEventEndpoint deletes an event
func DeleteEventEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	eventID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	deletedEvent := service.DeleteEvent(eventID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedEvent)
}

// UpdateEventEndpoint updates an event
func UpdateEventEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	eventID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var event model.Event
	_ = json.NewDecoder(r.Body).Decode(&event)
	updatedEvent := service.UpdateEvent(event, eventID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}
