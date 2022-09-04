package snippetrepo

import (
	"errors"
	"snippetdemo/pkg/models"
)

type User models.User

func (r *Repo) InsertUser(user *User) error {
	record := r.DbClient.Create(&user)

	return record.Error
}

func (r *Repo) GetUser(username string) (*User, error) {
	var user User
	record := r.DbClient.Where("username = ?", username).First(&user)

	if record.Error != nil {
		return nil, errors.New("user not found")
	} else {
		return &user, nil
	}
}
