package mongo

import (
	"context"
	"log"

	"github.com/Software-Architecture-2019-2/sa-event-ms/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbURL string = "mongodb://" + util.MongoURL

const dbName = "event"
const collectionName = "event"

var database *mongo.Database

// Database gett return the Database object.
func Database() *mongo.Database {
	connectoToDB()
	return database
}

// Connect establish a connection to database
func connectoToDB() {
	if database != nil {
		return
	}

	log.Println("Connecting to MongoDB " + dbURL)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURL))

	if err != nil {
		log.Fatal(err)
	}
	database = client.Database(dbName)
	log.Println("Connected to DB " + dbName)
}
