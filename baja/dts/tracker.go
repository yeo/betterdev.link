package dts

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

type TrackerService struct {
	DB *mongo.Database
}

func (t *TrackerService) Track(action string, payload map[string]interface{}) {

}

func (t *TrackerService) LoadIssue(issue, email, ip string) error {
	now := time.Now()
	secs := now.Unix()

	result, err := t.DB.Collection("activity").InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("action", "open_issue"),
			bson.EC.String("ip", ip),
			bson.EC.DateTime("created_at", secs*1000),
			bson.EC.SubDocumentFromElements("props",
				bson.EC.String("user", email),
				bson.EC.String("issue", issue),
			)))
	if err != nil {
		log.Println("Fail to save to mongodb", err)
		return err
	} else {
		log.Debug("Succesfully save", result)
		return nil
	}
}

func (t *TrackerService) OpenURL(link, email, ip string) error {
	result, err := t.DB.Collection("activity").InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("action", "open_url"),
			bson.EC.String("ip", ip),
			bson.EC.SubDocumentFromElements("props",
				bson.EC.String("url", link),
				bson.EC.String("user", email)),
		))

	if err != nil {
		log.Println("Fail to save to mongodb", err)
		return err
	}
	log.Debug("Succesfully save", result)

	return nil
}
