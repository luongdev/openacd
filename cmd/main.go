package main

import (
	"fmt"
	"github.com/luongdev/openacd/types"
)

func main() {
	//conf := config.LoadConfig()
	//
	//client, err := database.Connect(conf.Database)
	//if err != nil {
	//	panic(err)
	//}
	//
	//collection, err := client.Database("openacd").Collection("users")
	//if err != nil {
	//	return
	//}
	//
	//one, err := collection.InsertOne(context.Background(), bson.M{
	//	"_id":  bson.NewObjectID(),
	//	"name": "a",
	//	"age":  20,
	//})
	//if err != nil {
	//	return
	//}
	//
	//fmt.Printf("inserted id: %s\n", one.InsertedID)
	//
	//defer func() {
	//	_ = client.Disconnect()
	//}()

	f := types.NewCriterionFactory()
	c, err := f.New(
		types.WithName("ready"),
		types.WithDisplayName("Ready"),
		types.WithScore(100),
		types.WithWeight(1),
		types.WithMaxScore(100),
		types.WithType(types.StatusCriterion),
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("criterion score: %v\n", c.CalculateScore())
}
