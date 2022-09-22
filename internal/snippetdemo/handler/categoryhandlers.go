package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snippetdemo/internal/snippetdemo/service"
	"snippetdemo/pkg/models"
)

type CategoryHandler struct {
	svc service.CategoryService
}

func NewCategoryHandler(svc service.CategoryService) *CategoryHandler {
	h := CategoryHandler{svc: svc}
	return &h
}

func (h *CategoryHandler) FilterCategories(w http.ResponseWriter, r *http.Request) {
	type CategoryFilterReq struct {
		CategoryId    *int
		Description   *string
		SortBy        *string
		SortDirection *string
		Count         *int
	}

	var req CategoryFilterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.svc.GetCategories(models.CategoryFilter{
		CategoryId:    req.CategoryId,
		Description:   req.Description,
		SortBy:        req.SortBy,
		SortDirection: req.SortDirection,
		Count:         req.Count,
	})

	fmt.Println(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		render(w, res, http.StatusCreated)
	}

}
