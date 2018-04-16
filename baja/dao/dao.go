package dao

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	//"github.com/stretchr/testify/require"
	"log"
)

type TrackDAO struct {
	Server   string
	Database string
}

func Connect(URI string) *mongo.Database {
	client, err := mongo.NewClient(URI)
	log.Println("e", err)
	//require.NoError(client, err)

	db := client.Database("bd")
	return db
}

func (m *TrackDAO) Insert() {

}
