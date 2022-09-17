package main

import (
	"snippetdemo/internal/snippetdemo/handler"
	"snippetdemo/internal/snippetdemo/service"
)

func (srv *Server) MapHandlers() {
	snippetservice := service.NewSnippetService(srv.client)
	snippetHandler := handler.NewSnippetHandler(*snippetservice)

	// userservice := service.NewUserService(repo)

	srv.router.HandleFunc("/create-snippet", snippetHandler.CreateSnippet)
}
