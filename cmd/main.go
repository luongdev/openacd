package main

import (
	"context"
	"fmt"
	"github.com/luongdev/openacd/config"
	"github.com/luongdev/openacd/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	conf := config.LoadConfig()

	client, err := database.Connect(conf.Database)
	if err != nil {
		panic(err)
	}

	collection, err := client.Database("openacd").Collection("users")
	if err != nil {
		return
	}

	one, err := collection.InsertOne(context.Background(), bson.M{
		"_id":  bson.NewObjectID(),
		"name": "a",
		"age":  20,
	})
	if err != nil {
		return
	}

	fmt.Printf("inserted id: %s\n", one.InsertedID)

	defer func() {
		_ = client.Disconnect()
	}()
}
