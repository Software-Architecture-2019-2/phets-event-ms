package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Software-Architecture-2019-2/sa-event-ms/model"
	"github.com/Software-Architecture-2019-2/sa-event-ms/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const eventCollectionName = "event"

var eventCollection = mongo.Database().Collection(eventCollectionName)

// InsertOneValue inserts one item from Event model
func InsertOneValue(event model.Event) {
	inserted, err := eventCollection.InsertOne(context.Background(), event)
	if err != nil {
		log.Fatal(err)
	}
	println(fmt.Sprintf("%v", inserted.InsertedID))
	// return inserted.InsertedID
}

// GetAllEvents returns all events from DB
func GetAllEvents() []model.Event {
	cursor, err := eventCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var events []model.Event
	var event model.Event
	// Get the next result from the cursor
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&event)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.Background())
	return events
}

// GetEvent returns event from DB using event ID
func GetEvent(eventID primitive.ObjectID) model.Event {
	filter := bson.M{
		"_id": eventID,
	}
	var event model.Event
	err := eventCollection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		log.Fatal(err)
	}
	return event
}

// DeleteEvent deletes an existing event and returns the deleted event
func DeleteEvent(eventID primitive.ObjectID) model.Event {
	var deletedEvent model.Event
	filter := bson.M{
		"_id": eventID,
	}
	err := eventCollection.FindOneAndDelete(
		context.Background(),
		filter).Decode(&deletedEvent)
	if err != nil {
		log.Fatal(err)
	}
	return deletedEvent
}

// UpdateEvent updates an existing event and returns the update event
func UpdateEvent(event model.Event, eventID primitive.ObjectID) model.Event {
	filter := bson.M{
		"_id": eventID,
	}
	update := bson.M{
		"$set": bson.M{
			"subject":     event.Subject,
			"description": event.Description,
			"animal":      event.AnimalID,
			"date":        event.Date,
		},
	}

	var updatedEvent model.Event
	err := eventCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update).Decode(&updatedEvent)

	if err != nil {
		log.Fatal(err)
	}
	return updatedEvent
}
