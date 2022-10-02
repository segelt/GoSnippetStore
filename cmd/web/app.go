package main

import (
	"net/http"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	client    *mongo.Client
	router    *http.ServeMux
	secretKey string
	port      string
}

func (app *App) NewServer() *http.Server {
	app.router = http.NewServeMux()
	app.MapHandlers()

	//** CORS SPECIFIC SECTION **//
	allowedRouter := cors.AllowAll().Handler(app.router)

	server := &http.Server{
		Addr:    app.port,
		Handler: allowedRouter,
	}

	return server
}
