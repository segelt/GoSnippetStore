package main

import (
	"net/http"

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

	err := http.ListenAndServe(":4000", srv.router)
	return err
}
