package service

import (
	snippetrepo "snippetdemo/internal/snippetdemo/repo/postgres"
	"snippetdemo/pkg/models"
)

var err error

type Snippet models.Snippet

type SnippetService struct {
	Repo snippetrepo.Repo
}

func (svc *SnippetService) InsertSnippet(userId int, content string) error {
	err := svc.Repo.InsertSnippet(userId, content)

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

func NewSnippetService(repo snippetrepo.Repo) *SnippetService {
	return &SnippetService{Repo: repo}
}
