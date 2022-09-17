package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"userID" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password []byte             `json:"hashedPassword" bson:"password"`
}

func (user User) IdString() string {
	return user.ID.Hex()
}

type UserService interface {
	RegisterUser(username string, password string) error
	VerifyUser(username string, password string) (bool, error)
}
