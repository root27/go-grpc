package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Userdb *mongo.Collection
var db *mongo.Client

func Connect() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	db, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	Userdb = db.Database("go-grpc").Collection("users")

	return Userdb, nil
}
