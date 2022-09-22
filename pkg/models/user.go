package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"userID" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"hashedPassword" bson:"password"`
}

func (user User) IdString() string {
	return user.ID.Hex()
}

type UserService interface {
	RegisterUser(username string, password string) error
	VerifyUser(username string, password string) (bool, error)
}

type UserModel struct {
	Client *mongo.Client
}

func (u *UserModel) Get(userId string) (*User, error) {
	userscol := u.Client.Database("snippetdb").Collection("users")

	var user User
	err := userscol.FindOne(context.TODO(), bson.M{"userId": userId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
