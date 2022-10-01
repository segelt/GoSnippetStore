package handler

import (
	"encoding/json"
	"net/http"
	"snippetdemo/internal/snippetdemo/service"
	"snippetdemo/pkg/models"
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
		Content    string
		Title      string
		CategoryId int
	}

	var req CreateSnippetReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid := r.Context().Value("userid").(string)
	err = h.svc.InsertSnippet(userid, req.Content, req.Title, req.CategoryId)

	if err != nil {
		render(w, err, http.StatusInternalServerError)
	} else {
		render(w, "Success", http.StatusCreated)
	}
}

func (h *SnippetHandler) ViewSnippets(w http.ResponseWriter, r *http.Request) {
	type ViewSnippetReq struct {
		SortBy        *string
		SortDirection *string
		Page          *int
		PageSize      *int
	}

	var req ViewSnippetReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid := r.Context().Value("userid").(string)
	snippets, err := h.svc.GetSnippetsOfUser(models.SnippetFilter{
		UserId:        &userid,
		SortBy:        req.SortBy,
		SortDirection: req.SortDirection,
		PageSize:      req.PageSize,
		Page:          req.Page,
	})

	if err != nil {
		render(w, err, http.StatusInternalServerError)
	} else {
		render(w, snippets, http.StatusOK)
	}
}
