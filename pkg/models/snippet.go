package models

import (
	"context"
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

func (s *SnippetModel) Single(snippetId string) (Snippet, error) {
	panic("Not implemented..")
}

func (s *SnippetModel) Insert(userid string, content string, title string, categoryId int) error {
	panic("Not implemented..")
}

func (s *SnippetModel) Delete(snippetId string) (bool, error) {
	panic("Not implemented..")
}
