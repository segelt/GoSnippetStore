package main

import (
	"net/http"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	client    *mongo.Client
	router    *http.ServeMux
	secretKey string
}

func (srv *Server) StartServer() error {
	srv.router = http.NewServeMux()
	srv.MapHandlers()

	//** CORS SPECIFIC SECTION **//
	allowedRouter := cors.AllowAll().Handler(srv.router)
	err := http.ListenAndServe(":4000", allowedRouter)
	// ** ** //
	// err := http.ListenAndServe(":4000", srv.router)

	return err
}
