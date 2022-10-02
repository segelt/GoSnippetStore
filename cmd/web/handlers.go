package main

import (
	"snippetdemo/internal/middlewares"
	"snippetdemo/internal/snippetdemo/handler"
	"snippetdemo/internal/snippetdemo/service"
)

func (srv *App) MapHandlers() {
	dbname := "snippetdb"
	snippetservice := service.NewSnippetService(srv.client, dbname)
	snippetHandler := handler.NewSnippetHandler(*snippetservice)

	userservice := service.NewUserService(srv.client, srv.secretKey, dbname)
	userHandler := handler.NewUserHandler(*userservice)

	categoryService := service.NewCategoryService(srv.client, dbname)
	categoryHandler := handler.NewCategoryHandler(*categoryService)

	middlewareManager := middlewares.MiddleWareManager{SecretKey: srv.secretKey}

	srv.router.HandleFunc("/create-snippet", middlewares.MultipleMiddleware(snippetHandler.CreateSnippet, middlewareManager.Auth))
	srv.router.HandleFunc("/snippets", middlewares.MultipleMiddleware(snippetHandler.ViewSnippets, middlewareManager.Auth))
	srv.router.HandleFunc("/snippet", middlewares.MultipleMiddleware(snippetHandler.GetSnippet, middlewareManager.Auth))
	srv.router.HandleFunc("/createuser", userHandler.RegisterUser)
	srv.router.HandleFunc("/verifyuser", userHandler.VerifyUser)
	srv.router.HandleFunc("/categories", categoryHandler.FilterCategories)
}
