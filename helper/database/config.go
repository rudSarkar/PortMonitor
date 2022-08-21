package database

import (
	"context"
	"github.com/thep0y/go-logger/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IsConnected(URI string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal("Could not connect to database.")
	}

	return client
}
