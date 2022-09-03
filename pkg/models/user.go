package models

type User struct {
	ID       int    `json:"userID"`
	Username string `json:"username"`
	Password []byte `json:"hashedPassword"`
}

type UserService interface {
	RegisterUser(username string, password string) (int, error)
	VerifyUser(username string, password string) (bool, error)
}
