package mongoclient

import (
	"github.com/go-bongo/bongo"
	"github.com/sirupsen/logrus"
)

type MongoDB struct {
	connection *bongo.Connection
	Collection *bongo.Collection
	// session    *mgo.Session
}

func NewCLient(connection string, database string, collection string) *MongoDB {
	logrus.Println("Mongo DB - Client Connecting")
	logrus.Println("Mongo DB - connection string: ", connection)
	logrus.Println("Mongo DB - db name: ", database)

	config := &bongo.Config{
		ConnectionString: connection,
		Database:         database,
	}

	bongoConnection, err := bongo.Connect(config)
	if err != nil {
		logrus.Errorln("Mongo DB - Lib: mgo")
		panic(err)
	}

	coll := bongoConnection.Collection(collection)

	// session, err := mgo.Dial(connection)
	// if err != nil {
	// 	logrus.Errorln("Mongo DB - Lib: mgo")
	// 	panic(err)
	// }

	// col1 := session.DB(database).C(collection)

	logrus.Printf("Mongo Db - Connected")

	return &MongoDB{
		connection: bongoConnection,
		Collection: coll,
	}
	// return &MongoDB{
	// 	session: session,
	// }
}

func (m *MongoDB) Close() {
	logrus.Print("Closing Mongo Client Connection ... ")
	m.connection.Session.Close()
	// m.session.Close()
	logrus.Println("Closing Mongo Client Success")
}

func (m *MongoDB) GetCollection() *bongo.Collection {
	return m.Collection
}
