package mongoclient

import (
	"github.com/go-bongo/bongo"
	"github.com/sirupsen/logrus"
)

type MongoDB struct {
	connection *bongo.Connection
	Collection *bongo.Collection
}

func NewCLient(connection string, database string, collection string) *MongoDB {
	logrus.Println("Mongo DB - Client Connecting")
	logrus.Println("Mongo DB - conncetion string: ", connection)

	config := &bongo.Config{
		ConnectionString: connection,
		Database:         database,
	}

	bongoConnection, err := bongo.Connect(config)
	if err != nil {
		panic(err)
	}

	coll := bongoConnection.Collection(collection)

	logrus.Printf("Mongo Db - Connected")

	return &MongoDB{
		connection: bongoConnection,
		Collection: coll,
	}
}

func (m *MongoDB) Close() {
	logrus.Print("Closing Mongo Client Connection ... ")
	m.connection.Session.Close()
	logrus.Println("Closing Mongo Client Success")
}

func (m *MongoDB) GetCollection() *bongo.Collection {
	return m.Collection
}
