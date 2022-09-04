package main

import (
	"snippetdemo/internal/snippetdemo/handler"
	snippetrepo "snippetdemo/internal/snippetdemo/repo/postgres"
	"snippetdemo/internal/snippetdemo/service"
)

func (srv *Server) MapHandlers() {
	repo := snippetrepo.Repo{DbClient: srv.db}
	snippetservice := service.NewSnippetService(repo)
	snippetHandler := handler.NewSnippetHandler(*snippetservice)

	// userservice := service.NewUserService(repo)

	srv.router.HandleFunc("/create-snippet", snippetHandler.CreateSnippet)
}
