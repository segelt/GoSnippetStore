package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5433
	user     = "snippetdemouser"
	password = "snippetdemopwd1234"
	dbname   = "snippetdemodb"
)

var (
	DbClient *gorm.DB
)

func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DbClient, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	// defer dbClient.Close()
	fmt.Println((DbClient))
	fmt.Println("Successfully connected!")
}
