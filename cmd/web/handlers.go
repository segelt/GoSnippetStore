package main

import (
	"snippetdemo/internal/middlewares"
	"snippetdemo/internal/snippetdemo/handler"
	"snippetdemo/internal/snippetdemo/service"
)

func (srv *Server) MapHandlers() {
	snippetservice := service.NewSnippetService(srv.client)
	snippetHandler := handler.NewSnippetHandler(*snippetservice)

	userservice := service.NewUserService(srv.client, srv.secretKey)
	userHandler := handler.NewUserHandler(*userservice)

	middlewareManager := middlewares.MiddleWareManager{SecretKey: srv.secretKey}

	srv.router.HandleFunc("/create-snippet", middlewares.MultipleMiddleware(snippetHandler.CreateSnippet, middlewareManager.Auth))
	srv.router.HandleFunc("/createuser", userHandler.RegisterUser)
	srv.router.HandleFunc("/verifyuser", userHandler.VerifyUser)
}
