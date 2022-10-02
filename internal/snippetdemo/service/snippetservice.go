package service

import (
	"snippetdemo/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type SnippetService struct {
	snippets models.SnippetModel
}

// type InsertSnippetReq struct {
// 	UserId     string
// 	Content    string
// 	Title      string
// 	CategoryId *int
// }

func (svc *SnippetService) InsertSnippet(userId string, content string, title string, categoryId int) error {

	err := svc.snippets.Insert(userId, content, title, categoryId)
	return err
}
func (svc *SnippetService) GetSnippetById(snippetId string) (*models.Snippet, error) {
	snippet, err := svc.snippets.Single(snippetId)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}
func (svc *SnippetService) GetSnippetsOfUser(filter models.SnippetFilter) (*[]models.Snippet, error) {

	if filter.PageSize == nil {
		var defaultPageSizeInt int = 50
		filter.PageSize = &defaultPageSizeInt
	}

	if filter.Page == nil {
		var defaultPage int = 0
		filter.PageSize = &defaultPage
	}

	snippets, err := svc.snippets.ByUser(filter)
	if err != nil {
		return nil, err
	}

	return &snippets, nil
}
func (svc *SnippetService) DeleteSnippet(snippetId string) (bool, error) {
	res, err := svc.snippets.Delete(snippetId)
	return res, err
}

func NewSnippetService(client *mongo.Client, DBName string) *SnippetService {
	return &SnippetService{snippets: models.SnippetModel{Client: client, DBName: DBName}}
}
