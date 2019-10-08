package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event asdasd
type Event struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Subject     string             `bson:"subject" json:"subject"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date"`
	AnimalID    string             `bson:"animal_id" json:"animal_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
