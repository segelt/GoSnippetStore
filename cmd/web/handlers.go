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

	categoryService := service.NewCategoryService(srv.client)
	categoryHandler := handler.NewCategoryHandler(*categoryService)

	middlewareManager := middlewares.MiddleWareManager{SecretKey: srv.secretKey}

	srv.router.HandleFunc("/create-snippet", middlewares.MultipleMiddleware(snippetHandler.CreateSnippet, middlewareManager.Auth))
	srv.router.HandleFunc("/snippets", middlewares.MultipleMiddleware(snippetHandler.ViewSnippets, middlewareManager.Auth))
	srv.router.HandleFunc("/createuser", userHandler.RegisterUser)
	srv.router.HandleFunc("/verifyuser", userHandler.VerifyUser)
	srv.router.HandleFunc("/categories", categoryHandler.FilterCategories)
}
