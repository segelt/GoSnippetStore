package service

import (
	"context"
	"errors"
	"snippetdemo/internal/snippetdemo/helpers"
	"snippetdemo/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Client    *mongo.Client
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
	coll := svc.Client.Database("snippetdb").Collection("users")
	userd := bson.D{{Key: "username", Value: username}, {Key: "password", Value: hashedpwd}}
	_, err := coll.InsertOne(context.TODO(), userd)
	return err
}
func (svc *UserService) VerifyUser(username string, password string) (string, error) {
	db := svc.Client.Database("snippetdb")
	userscol := db.Collection("users")

	var user models.User
	err := userscol.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		return "", err
	}

	err = CheckPassword(user.Password, password)
	if err != nil {
		return "", err
	}

	// User is validated, generate JWT token
	generatedToken, err := svc.generateToken(user.IdString(), user.Username, svc.Secretkey)

	if err != nil {
		return "", err
	}

	return generatedToken, nil
}

func (svc *UserService) generateToken(userId string, username string, secretkey string) (string, error) {
	mgr := helpers.JWTManager{Secretkey: []byte(svc.Secretkey)}

	return mgr.GenerateJWT(userId, username)
}

func NewUserService(client *mongo.Client, secretkey string) *UserService {
	return &UserService{Client: client, Secretkey: secretkey}
}
