package main

import (
	"net/http"
)

func snippets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet demo test"))
}
