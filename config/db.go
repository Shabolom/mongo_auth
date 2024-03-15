package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func ConnectionDB() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Env.DbUri))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	var db = client.Database(Env.DbName)
	var collection = db.Collection(Env.TableName)

	Collection = collection

	fmt.Println("Connected to MongoDB!")
	return nil
}
