package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Snippet struct {
	ID       primitive.ObjectID `json:"snippetID" bson:"_id,omitempty"`
	UserID   int                `json:"userID" bson:"userid"`
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
