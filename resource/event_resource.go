package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Software-Architecture-2019-2/phets-event-ms/model"
	"github.com/Software-Architecture-2019-2/phets-event-ms/service"
	"github.com/Software-Architecture-2019-2/phets-event-ms/util"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var people []model.Event

// GetEventByID gets an event
func GetEventByID(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	eventID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		errorResponse := util.GenerateErrorResponse(500, "Unexpected error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	event, err := service.GetEvent(eventID)
	if err != nil {
		errorResponse := util.GenerateErrorResponse(404, "Id not found", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

// GetAllEvents gets all events
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	events, err := service.GetAllEvents()
	if err != nil {
		errorResponse := util.GenerateErrorResponse(500, "Unexpected error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

// CreateEvent creates an event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var event model.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse := util.GenerateErrorResponse(400, "Not body found in the request.", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = json.Unmarshal(reqBody, &event)
	if err != nil {
		errorResponse := util.GenerateErrorResponse(400, "The inserted fields does not coincide.", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	insertedEvent, err := service.InsertOneValue(event)
	if err != nil {
		errorResponse := util.GenerateErrorResponse(500, "An error occurred while processing request.", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedEvent)
}

// DeleteEvent deletes an event
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	eventID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		errorResponse := util.GenerateErrorResponse(500, "Unexpected error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	deletedEvent, err := service.DeleteEvent(eventID)
	if err != nil {
		errorResponse := util.GenerateErrorResponse(404, "Id not found", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedEvent)
}

// UpdateEvent updates an event
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	eventID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		errorResponse := util.GenerateErrorResponse(500, "Unexpected error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var event model.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		errorResponse := util.GenerateErrorResponse(406, "Params are not valid", err)
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	updatedEvent, err := service.UpdateEvent(event, eventID)

	if err != nil {
		errorResponse := util.GenerateErrorResponse(404, "Id not found", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}
