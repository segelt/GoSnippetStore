package service

import (
	"errors"
	"snippetdemo/internal/snippetdemo/helpers"
	snippetrepo "snippetdemo/internal/snippetdemo/repo/postgres"
	"snippetdemo/pkg/models"
)

type User models.User

type UserService struct {
	Repo snippetrepo.Repo
}

func HashPassword(password string) string {
	hashedPassword := helpers.HashStr(password)
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

func (svc *UserService) RegisterUser(username string, password string) (int, error) {
	panic("Not implemented")
}
func (svc *UserService) VerifyUser(username string, password string) (bool, error) {
	panic("Not implemented")
}
