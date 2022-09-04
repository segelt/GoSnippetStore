package snippetrepo

import (
	"snippetdemo/pkg/models"
	"time"
)

type Snippet models.Snippet

func (r *Repo) InsertSnippet(userId int, content string) error {
	expireTime := time.Now().AddDate(0, 0, 10)
	result := r.DbClient.Create(&Snippet{Content: content, UserID: userId, Created: time.Now(), Expires: expireTime})

	return result.Error
}

// func GetSnippetById(id int) (*Snippet, error) {
// 	panic("Not implemented.")
// }
// func GetSnippetsOfUser(userId int) ([]*Snippet, error) {
// 	panic("Not implemented.")
// }
// func DeleteSnippet(id int) (bool, error) {
// 	panic("Not implemented.")
// }
