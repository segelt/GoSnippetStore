package models

import (
	"context"
	"errors"
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
	RegisterUser(ctx context.Context, username string, password string) error
	VerifyUser(ctx context.Context, username string, password string) (bool, error)
}

type UserModel struct {
	Client *mongo.Client
	DBName string
}

type UserFilter struct {
	Username *string
	// Email    *string
}

func (u *UserModel) Get(ctx context.Context, userId string) (*User, error) {
	userscol := u.Client.Database(u.DBName).Collection("users")

	var user User
	err := userscol.FindOne(ctx, bson.M{"_id": objectIDFromHex(userId)}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserModel) Filter(ctx context.Context, filter UserFilter) (*[]User, error) {
	coll := u.Client.Database(u.DBName).Collection("users")

	qry := bson.D{}
	if filter.Username != nil {
		f := bson.E{Key: "username",
			Value: bson.D{{Key: "$regex",
				Value: *filter.Username}}}
		qry = append(qry, f)
	}

	cursor, err := coll.Find(ctx, qry)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserModel) FilterSingle(ctx context.Context, filter UserFilter) (*User, error) {
	if filter.Username == nil {
		return nil, errors.New("username filter cannot be nil")
	}

	coll := u.Client.Database(u.DBName).Collection("users")

	var user User
	err := coll.FindOne(ctx, bson.D{{Key: "username", Value: *filter.Username}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserModel) Insert(ctx context.Context, userName string, password string) error {
	coll := u.Client.Database(u.DBName).Collection("users")

	qry := bson.D{}
	f := bson.E{Key: "username",
		Value: bson.D{{Key: "$regex",
			Value: primitive.Regex{
				Pattern: userName,
				Options: "i"}},
		},
	}
	qry = append(qry, f)

	var user User
	err := coll.FindOne(ctx, qry).Decode(&user)

	if err == nil {
		return errors.New("user already exists")
	}

	userd := bson.D{{Key: "username", Value: userName}, {Key: "password", Value: password}}
	_, err = coll.InsertOne(ctx, userd)
	return err
}
