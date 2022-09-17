package main

import (
	"fmt"
	"log"
	"os"
	"snippetdemo/internal/database/mongorepo"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUri := os.Getenv("MONGODB_URI")

	s := mongorepo.Repo{}

	err = s.Initialize(mongoUri)

	fmt.Println("Started server")

	if err != nil {
		log.Fatal(err)
	}

}
