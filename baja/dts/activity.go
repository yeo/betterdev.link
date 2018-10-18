package dts

import (
	"context"
	//"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

type ActivityService struct {
	DB *mongo.Database
}

func (t *ActivityService) LoadUser(email string) []TrackerAction {
	var result []TrackerAction

	ctx := context.Background()
	cursor, err := t.DB.Collection("activity").Find(ctx,
		bson.NewDocument(
			bson.EC.String("props.user", email),
		))
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem TrackerAction
		if err := cursor.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		log.Println(elem)

		result = append(result, elem)
	}

	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
