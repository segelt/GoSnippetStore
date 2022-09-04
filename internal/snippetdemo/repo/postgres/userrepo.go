package snippetrepo

import (
	"errors"
	"snippetdemo/pkg/models"
)

func (r *Repo) InsertUser(user *models.User) error {
	record := r.DbClient.Create(&user)

	return record.Error
}

func (r *Repo) GetUser(username string) (*models.User, error) {
	var user models.User
	record := r.DbClient.Where("username = ?", username).First(&user)

	if record.Error != nil {
		return nil, errors.New("user not found")
	} else {
		return &user, nil
	}
}
