package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Software-Architecture-2019-2/phets-event-ms/model"
	"github.com/Software-Architecture-2019-2/phets-event-ms/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const eventCollectionName = "event"

var eventCollection = mongo.Database().Collection(eventCollectionName)

// InsertOneValue inserts one item from Event model
func InsertOneValue(event model.Event) model.Event {
	event.ID = primitive.NewObjectID()
	_, err := eventCollection.InsertOne(context.Background(), event)
	if err != nil {
		log.Fatal(err)
	}

	return event
}

// GetAllEvents returns all events from DB
func GetAllEvents() []model.Event {
	cursor, _ := eventCollection.Find(context.Background(), bson.D{})

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
		"_id": bson.M{
			"$eq": eventID,
		},
	}
	var event model.Event
	eventCollection.FindOne(context.Background(), filter).Decode(&event)
	return event
}

// DeleteEvent deletes an existing event and returns the deleted event
func DeleteEvent(eventID primitive.ObjectID) model.Event {
	var deletedEvent model.Event
	filter := bson.M{
		"_id": bson.M{
			"$eq": eventID,
		},
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
		"_id": bson.M{
			"$eq": eventID,
		},
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
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEvent)

	if err != nil {
		log.Fatal(err)
	}
	return updatedEvent
}
