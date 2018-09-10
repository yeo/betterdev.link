package dts

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

type TrackerProp struct {
	User  string `json:"user"`
	Issue string `json:"issue"`
}

type TrackerAction struct {
	ID        string      `json:"_id"`
	Action    string      `json:"action"`
	CreatedAt time.Time   `json:"created_at"`
	Props     TrackerProp `json:"props"`
}

type TrackerService struct {
	DB *mongo.Database
}

func (t *TrackerService) AlreadySent(issue, email string) bool {
	var result TrackerAction

	err := t.DB.Collection("activity").FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("action", "sent_issue"),
			bson.EC.SubDocumentFromElements("props",
				bson.EC.String("user", email),
				bson.EC.String("issue", issue),
			))).Decode(&result)

	log.Println("Check sent for", issue, email)
	if err != nil {
		log.Println("error", err)
		return false
	}

	log.Println("Result", result, result.ID)

	return result.Props.User == email
}

func (t *TrackerService) LogSent(issue, email string) error {
	now := time.Now()
	secs := now.Unix()

	result, err := t.DB.Collection("activity").InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("action", "sent_issue"),
			bson.EC.DateTime("created_at", secs*1000),
			bson.EC.SubDocumentFromElements("props",
				bson.EC.String("user", email),
				bson.EC.String("issue", issue),
			)))
	if err != nil {
		log.Println("Cannot mark sent status for issue", issue, "email", email, err)
		return err
	} else {
		log.Debug("Succesfully mark sent", result)
		return nil
	}
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

func (t *TrackerService) OpenURL(link, issue, email, ip string) error {
	result, err := t.DB.Collection("activity").InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("action", "open_url"),
			bson.EC.String("ip", ip),
			bson.EC.SubDocumentFromElements("props",
				bson.EC.String("issue", issue),
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
