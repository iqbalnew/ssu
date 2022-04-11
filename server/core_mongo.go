package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongo_client *mongo.Client
)

func startMongoConnection() {
	logrus.Printf("Starting Mongo Db Connections...")

	initMongoMain()

}

func closeMongoConnections() {
	closeMongoMain()
}

func initMongoMain() {
	logrus.Printf("Mongo Db - Connecting")

	mongo_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	if err := mongo_client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	logrus.Printf("Mongo Db - Connected")
}

func closeMongoMain() {
	logrus.Print("Closing DB Main Connection ... ")
	if err := mongo_client.Disconnect(context.TODO()); err != nil {
		logrus.Fatalf("Error on disconnection with DB Main : %v", err)
	}
	logrus.Println("Closing DB Main Success")
}
