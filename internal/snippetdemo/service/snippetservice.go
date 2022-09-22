package service

import (
	"context"
	"fmt"
	"log"
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

// type InsertSnippetReq struct {
// 	UserId     string
// 	Content    string
// 	Title      string
// 	CategoryId *int
// }

func (svc *SnippetService) InsertSnippet(userId string, content string, title string, categoryid int) error {

	categorycl := svc.Client.Database("snippetdb").Collection("categories")
	var cg models.Category

	err = categorycl.FindOne(context.TODO(), bson.D{{"categoryId", categoryid}}).Decode(&cg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No categories match this query. %d\n", categoryid)
			return fmt.Errorf("No categories match this query. %d", categoryid)
		}
		return err
	}

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
