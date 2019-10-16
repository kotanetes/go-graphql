package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbSession *mongo.Client

func OpenConnection() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?gssapiServiceName=mongodb")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("error connecting to DB", err)
	}

	DbSession = client
}
