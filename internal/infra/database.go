package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBDatabase struct {
	client *mongo.Client
}
const URI = "mongodb://localhost:27017"

func ConnectDB() (*MongoDBDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	db := &MongoDBDatabase{
		client: client,
	}
	return db, nil
}

func (d *MongoDBDatabase) Close() error {
	return d.client.Disconnect(context.Background())
}


func (d *MongoDBDatabase) GetClient() *mongo.Client {
	return d.client
}