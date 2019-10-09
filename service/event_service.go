package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Software-Architecture-2019-2/phets-event-ms/model"
	"github.com/Software-Architecture-2019-2/phets-event-ms/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const eventCollectionName = "event"

var eventCollection = mongo.Database().Collection(eventCollectionName)

// InsertOneValue inserts one item from Event model
func InsertOneValue(event model.Event) (model.Event, error) {
	event.ID = primitive.NewObjectID()
	event.CreatedAt = time.Now()

	_, err := eventCollection.InsertOne(context.Background(), event)

	return event, err
}

// GetAllEvents returns all events from DB
func GetAllEvents() (model.EventsList, error) {
	cursor, _ := eventCollection.Find(context.Background(), bson.D{})

	events := model.EventsList{
		Total: 0,
		List:  []model.Event{},
	}
	var event model.Event
	var err error
	// Get the next result from the cursor
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&event)
		if err != nil {
			break
		}
		events.Total++
		events.List = append(events.List, event)
	}
	if err = cursor.Err(); err != nil {
		return events, err
	}
	cursor.Close(context.Background())
	return events, err
}

// GetEvent returns event from DB using event ID
func GetEvent(eventID primitive.ObjectID) (model.Event, error) {
	filter := bson.M{
		"_id": bson.M{
			"$eq": eventID,
		},
	}
	var event model.Event
	err := eventCollection.FindOne(context.Background(), filter).Decode(&event)
	return event, err
}

// DeleteEvent deletes an existing event and returns the deleted event
func DeleteEvent(eventID primitive.ObjectID) (model.Event, error) {
	var deletedEvent model.Event
	filter := bson.M{
		"_id": bson.M{
			"$eq": eventID,
		},
	}
	err := eventCollection.FindOneAndDelete(
		context.Background(),
		filter).Decode(&deletedEvent)
	return deletedEvent, err
}

// UpdateEvent updates an existing event and returns the update event
func UpdateEvent(event model.Event, eventID primitive.ObjectID) (model.Event, error) {
	filter := bson.M{
		"_id": bson.M{
			"$eq": eventID,
		},
	}
	update := bson.M{
		"$set": bson.M{
			"subject":     event.Subject,
			"description": event.Description,
			"animal_id":   event.AnimalID,
			"date":        event.Date,
			"updated_at":  time.Now(),
		},
	}

	var updatedEvent model.Event
	err := eventCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEvent)

	return updatedEvent, err
}
