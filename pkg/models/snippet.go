package models

import "time"

type Snippet struct {
	ID      int       `json:"snippetID"`
	UserID  int       `json:"userID"`
	Content string    `json:"snippetContent"`
	Created time.Time `json:"dateCreated"`
	Expires time.Time `json:"dateExpire"`
}

type SnippetService interface {
	InsertSnippet(userId int, content string) error
	GetSnippetById(id int) (*Snippet, error)
	GetSnippetsOfUser(userId int) ([]*Snippet, error)
	DeleteSnippet(id int) (bool, error)
}
