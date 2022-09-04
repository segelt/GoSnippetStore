package main

import (
	"fmt"
	"log"

	client "snippetdemo/internal/database/postgres"
	repo "snippetdemo/internal/snippetdemo/repo/postgres"
)

func main() {
	repo.MigrateModels()

	srv := &Server{db: client.DbClient}
	err := srv.StartServer()

	fmt.Println("Started server")

	if err != nil {
		log.Fatal(err)
	}

}
