package service

import (
	"context"
	"snippetdemo/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var err error

type Snippet models.Snippet

type SnippetService struct {
	Client *mongo.Client
}

func (svc *SnippetService) InsertSnippet(userId string, content string, title string, categoryid int) error {

	coll := svc.Client.Database("snippetdb").Collection("snippets")
	// err := svc.Repo.InsertSnippet(userId, content)
	createDate := time.Now()
	expireTime := createDate.AddDate(0, 0, 10)
	snippet := bson.D{{"content", content},
		{"UserId", userId},
		{"title", title},
		{"category", categoryid},
		{"created", createDate},
		{"expireDate", expireTime}}

	_, err := coll.InsertOne(context.TODO(), snippet)
	return err
}
func (svc *SnippetService) GetSnippetById(id int) (*Snippet, error) {
	panic("Not implemented")
}
func (svc *SnippetService) GetSnippetsOfUser(userId int) ([]*Snippet, error) {
	panic("Not implemented")
}
func (svc *SnippetService) DeleteSnippet(id int) (bool, error) {
	panic("Not implemented")
}

func NewSnippetService(client *mongo.Client) *SnippetService {
	return &SnippetService{Client: client}
}
