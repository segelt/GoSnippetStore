package handler

import (
	"encoding/json"
	"net/http"
	"snippetdemo/internal/snippetdemo/service"
)

type SnippetHandler struct {
	svc service.SnippetService
}

func NewSnippetHandler(svc service.SnippetService) *SnippetHandler {
	h := SnippetHandler{svc: svc}
	return &h
}

func (h *SnippetHandler) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	type CreateSnippetReq struct {
		Content string
	}

	var req CreateSnippetReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid := r.Context().Value("userid").(string)
	err = h.svc.InsertSnippet(userid, req.Content)

	if err != nil {
		render(w, err, http.StatusInternalServerError)
	}

	render(w, "Success", http.StatusCreated)
}
