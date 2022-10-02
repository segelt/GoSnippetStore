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
	log.SetOutput(os.Stdout)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUri := os.Getenv("MONGODB_URI")

	client, err := mongocl.NewMongoDB(mongoUri)
	if err != nil {
		log.Fatal(err)
	}

	err = mongocl.SeedData(client)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = mongocl.GracefulShutdownDbConnection(client)

		if err != nil {
			log.Fatal("Could not disconnect from db")
			log.Fatal(err)
			panic(err)
		}
	}()

	jwtkey := os.Getenv("JWT_KEY")
	srv := &Server{client: client, secretKey: jwtkey}
	err = srv.StartServer()
	fmt.Println("Started server")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
