package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Usuario struct {
	Name string `bson:"nome"`
	Age  int8   `bson:"idade"`
}

var (
	MONGO_URL     = "MONGO_URL"
	DATABASE_NAME = "DATABASE_NAME"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongo_uri := os.Getenv(MONGO_URL)
	database_name := os.Getenv(DATABASE_NAME)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(database_name), nil
}
