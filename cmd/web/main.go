package main

import (
	"log"
	"net/http"

	_ "snippetdemo/internal/database/postgres"
	repo "snippetdemo/internal/snippetdemo/repo/postgres"
)

func main() {
	repo.MigrateModels()

	mux := http.NewServeMux()
	mux.HandleFunc("/", snippets)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
