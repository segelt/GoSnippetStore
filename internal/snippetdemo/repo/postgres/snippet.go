package snippetrepo

import (
	"snippetdemo/pkg/models"
	"time"

	client "snippetdemo/internal/database/postgres"
)

type Snippet models.Snippet

type SnippetService models.SnippetService

func InsertSnippet(userId int, content string) (int, error) {
	expireTime := time.Now().Add(10)

	result := client.DbClient.Create(&Snippet{Content: content, UserID: userId, Created: time.Now(), Expires: expireTime})

	if result.Error != nil {
		panic("Error occured while inserting snippet")
	}

	return int(result.RowsAffected), nil
}
func GetSnippetById(id int) (*Snippet, error) {
	panic("Not implemented.")
}
func GetSnippetsOfUser(userId int) ([]*Snippet, error) {
	panic("Not implemented.")
}
func DeleteSnippet(id int) (bool, error) {
	panic("Not implemented.")
}
