package service

import (
	"errors"
	"snippetdemo/internal/snippetdemo/helpers"
	snippetrepo "snippetdemo/internal/snippetdemo/repo/postgres"
	"snippetdemo/pkg/models"
)

type UserService struct {
	Repo snippetrepo.Repo
}

func HashPassword(password string) []byte {
	hashedPassword := helpers.HashStrAsByteArray(password)
	return hashedPassword
}
func CheckPassword(actualHashedPassword string, providedPassword string) error {
	compareResult := helpers.CompareHashAndPassword(providedPassword, actualHashedPassword)
	if !compareResult {
		return errors.New("given password does not match for the current user")
	} else {
		return nil
	}
}

func (svc *UserService) RegisterUser(username string, password string) error {
	hashedpwd := HashPassword(password)
	user := models.User{Username: username, Password: hashedpwd}

	err := svc.Repo.InsertUser(&user)
	return err
}
func (svc *UserService) VerifyUser(username string, password string) (bool, error) {
	panic("Not implemented")
}

func NewUserService(repo snippetrepo.Repo) *UserService {
	return &UserService{Repo: repo}
}
