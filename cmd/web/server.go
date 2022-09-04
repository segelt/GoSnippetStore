package main

import (
	"net/http"

	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *http.ServeMux
}

func (srv *Server) StartServer() error {
	srv.router = http.NewServeMux()
	srv.MapHandlers()

	err := http.ListenAndServe(":4000", srv.router)
	return err
}
