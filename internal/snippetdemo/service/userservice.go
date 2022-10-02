package service

import (
	"errors"
	"snippetdemo/internal/snippetdemo/helpers"
	"snippetdemo/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Users     models.UserModel
	Secretkey string
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
	hashedpwd := helpers.HashStr(password)
	return svc.Users.Insert(username, hashedpwd)
}
func (svc *UserService) VerifyUser(username string, password string) (*string, error) {
	var user *models.User
	userFilter := models.UserFilter{Username: &username}
	user, err := svc.Users.FilterSingle(userFilter)

	if err != nil {
		return nil, err
	}

	err = CheckPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	// User is validated, generate JWT token
	generatedToken, err := svc.generateToken(user.IdString(), user.Username, svc.Secretkey)

	if err != nil {
		return nil, err
	}

	return &generatedToken, nil
}

func (svc *UserService) generateToken(userId string, username string, secretkey string) (string, error) {
	mgr := helpers.JWTManager{Secretkey: []byte(svc.Secretkey)}

	return mgr.GenerateJWT(userId, username)
}

func NewUserService(client *mongo.Client, secretkey string, DBName string) *UserService {
	return &UserService{Users: models.UserModel{Client: client, DBName: DBName}, Secretkey: secretkey}
}
