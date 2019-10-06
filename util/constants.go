package util

import (
	"os"
)

var mongoHost string = os.Getenv("MONGO_HOST")
var mongoPort string = os.Getenv("MONGO_PORT")

// MongoURL for DB connection.
var MongoURL string

func init() {
	if mongoHost == "" {
		mongoHost = "localhost"
	}

	if mongoPort == "" {
		mongoPort = "27017"
	}
	MongoURL = mongoHost + ":" + mongoPort
}
