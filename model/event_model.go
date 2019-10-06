package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event asdasd
type Event struct {
	ID          primitive.ObjectID `bson:"_id"`
	Subject     string             `bson:"subject"`
	Description string             `bson:"description"`
	Date        string             `bson:"date"`
	AnimalID    string             `bson:"animalId"`
}
