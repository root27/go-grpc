package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Client
}

func Connect() (*mongo.Client, error) {

	opt := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opt)

	if err != nil {
		return nil, err

	}

	return client, nil

}

func (c *Client) Database(name string) *mongo.Database {

	return c.Database(name)

}

func (c *Client) Collection(db, name string) *mongo.Collection {

	return c.Database(db).Collection(name)

}
