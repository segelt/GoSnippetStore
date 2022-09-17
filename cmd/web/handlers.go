package main

import (
	"snippetdemo/internal/snippetdemo/handler"
	"snippetdemo/internal/snippetdemo/service"
)

func (srv *Server) MapHandlers() {
	snippetservice := service.NewSnippetService(srv.client)
	snippetHandler := handler.NewSnippetHandler(*snippetservice)

	userservice := service.NewUserService(srv.client, srv.secretKey)
	userHandler := handler.NewUserHandler(*userservice)

	srv.router.HandleFunc("/create-snippet", snippetHandler.CreateSnippet)
	srv.router.HandleFunc("/createuser", userHandler.RegisterUser)
	srv.router.HandleFunc("/verifyuser", userHandler.VerifyUser)
}
