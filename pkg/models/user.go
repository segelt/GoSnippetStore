package models

import (
	"context"
	"log"

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

var objectIDFromHex = func(hex string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		log.Fatal(err)
	}
	return objectID
}

type UserService interface {
	RegisterUser(username string, password string) error
	VerifyUser(username string, password string) (bool, error)
}

type UserModel struct {
	Client *mongo.Client
}

type UserFilter struct {
	Username *string
	// Email    *string
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

func (u *UserModel) Filter(filter UserFilter) (*[]User, error) {
	coll := u.Client.Database("snippetdb").Collection("users")

	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "username", Value: *filter.Username}})
	if err != nil {
		return nil, err
	}

	var users []User
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return &users, nil
}
