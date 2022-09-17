package main

import (
	"fmt"
	"log"
	"os"
	"snippetdemo/internal/database/mongocl"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUri := os.Getenv("MONGODB_URI")
	s := mongocl.Repo{}
	err = s.Initialize(mongoUri)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	srv := &Server{client: s.Client}
	err = srv.StartServer()
	fmt.Println("Started server")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
