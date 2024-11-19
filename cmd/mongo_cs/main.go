package main

import (
	"context"
	"fmt"
	"mongo_cs/config"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Create aggregation pipeline
// TODO: Watch a change stream
// TODO: Run aggregation pipeline on docs from Change Streams
// TODO: Write results to stdout

func main() {
	fmt.Println(os.Getwd())
	conf := config.NewConfig()
	fmt.Println("Hello, World!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoUri := conf.MongoDBString
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	var result bson.M
	if err := db.Database("mongo_cs").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	col := db.Database("mongo_cs").Collection("inventory")
	cs, err := col.Watch(ctx, mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	defer cs.Close(ctx)
	ok := cs.Next(ctx)
	fmt.Println(ok)
	next := cs.Current
	fmt.Println(next)
}
