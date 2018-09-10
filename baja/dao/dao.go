package dao

import (
	"context"

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

	client.Connect(context.Background())
	db := client.Database("bd")
	return db
}

func (m *TrackDAO) Insert() {

}
