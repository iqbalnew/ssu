package mongoclient

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client *mongo.Client
}

func NewCLient(connection string) *MongoDB {
	logrus.Printf("Mongo Db - Client Connecting")

	mongo_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))
	if err != nil {
		panic(err)
	}

	if err := mongo_client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	logrus.Printf("Mongo Db - Connected")

	return &MongoDB{
		client: mongo_client,
	}
}

func (m *MongoDB) Close() {
	logrus.Print("Closing Mongo Client Connection ... ")
	if err := m.client.Disconnect(context.TODO()); err != nil {
		logrus.Fatalf("Error on disconnection closing Mongo Client : %v", err)
	}
	logrus.Println("Closing Mongo Client Success")
}
