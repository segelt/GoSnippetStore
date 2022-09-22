package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Snippet struct {
	ID       primitive.ObjectID `json:"snippetID" bson:"_id,omitempty"`
	UserID   int                `json:"userID" bson:"userId"`
	Category int                `json:"category" bson:"category"`
	Title    string             `json:"title" bson:"title"`
	Content  string             `json:"snippetContent" bson:"content"`
	Created  time.Time          `json:"dateCreated" bson:"created"`
	Expires  time.Time          `json:"dateExpire" bson:"expireDate"`
}

type SnippetService interface {
	InsertSnippet(userId string, content string, title string, categoryid int) error
	GetSnippetById(id int) (*Snippet, error)
	GetSnippetsOfUser(userId int) ([]*Snippet, error)
	DeleteSnippet(id int) (bool, error)
}

type SnippetModel struct {
	Client *mongo.Client
}

func (s *SnippetModel) ByUser(userId string) ([]Snippet, error) {
	coll := s.Client.Database("snippetdb").Collection("snippets")

	cursor, err := coll.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var targetSnippets []Snippet
	if err := cursor.All(context.TODO(), &targetSnippets); err != nil {
		return nil, err
	}

	return targetSnippets, nil
}

func (s *SnippetModel) Single(snippetId string) (*Snippet, error) {
	coll := s.Client.Database("snippetdb").Collection("snippets")

	var targetSnippet *Snippet
	err := coll.FindOne(context.TODO(), bson.M{"userId": snippetId}).Decode(targetSnippet)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return targetSnippet, nil
}

func (s *SnippetModel) Insert(userId string, content string, title string, categoryId int) error {

	categorycl := s.Client.Database("snippetdb").Collection("categories")
	var cg Category

	err := categorycl.FindOne(context.TODO(), bson.D{{Key: "categoryId", Value: categoryId}}).Decode(&cg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No categories match this query. %d\n", categoryId)
			return fmt.Errorf("no categories match this query. %d", categoryId)
		}
		return err
	}

	coll := s.Client.Database("snippetdb").Collection("snippets")
	// err := svc.Repo.InsertSnippet(userId, content)
	createDate := time.Now()
	expireTime := createDate.AddDate(0, 0, 10)
	snippet := bson.D{{Key: "content", Value: content},
		{Key: "userId", Value: userId},
		{Key: "title", Value: title},
		{Key: "category", Value: categoryId},
		{Key: "created", Value: createDate},
		{Key: "expireDate", Value: expireTime}}

	_, err = coll.InsertOne(context.TODO(), snippet)
	return err

}

func (s *SnippetModel) Delete(snippetId string) (bool, error) {

	categorycl := s.Client.Database("snippetdb").Collection("snippets")

	result, err := categorycl.DeleteOne(context.TODO(), bson.M{"_id": snippetId})
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, errors.New("no snippet was deleted")
	}

	return true, nil
}
