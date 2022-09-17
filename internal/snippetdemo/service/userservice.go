package service

import (
	"context"
	"errors"
	"snippetdemo/internal/snippetdemo/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Client *mongo.Client
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
	coll := svc.Client.Database("snippetdb").Collection("users")
	userd := bson.D{{"Username", username}, {"Password", hashedpwd}}
	_, err := coll.InsertOne(context.TODO(), userd)
	return err
}
func (svc *UserService) VerifyUser(username string, password string) (bool, error) {
	panic("Not implemented")
}

func NewUserService(client *mongo.Client) *UserService {
	return &UserService{Client: client}
}
