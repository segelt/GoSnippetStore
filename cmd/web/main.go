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

	defer func() {
		err = s.GracefulShutdownDbConnection()

		if err != nil {
			log.Fatal("Could not disconnect from db")
			log.Fatal(err)
			panic(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	jwtkey := os.Getenv("JWT_KEY")
	srv := &Server{client: s.Client, secretKey: jwtkey}
	err = srv.StartServer()
	fmt.Println("Started server")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
