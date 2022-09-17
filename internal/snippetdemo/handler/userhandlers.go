package handler

import (
	"encoding/json"
	"net/http"
	"snippetdemo/internal/snippetdemo/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	h := UserHandler{svc: svc}

	return &h
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	type UserRegisterReq struct {
		Username string
		Password string
	}

	var req UserRegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.svc.RegisterUser(req.Username, req.Password)

	if err != nil {
		render(w, err, http.StatusInternalServerError)
	}

	render(w, "Success", http.StatusCreated)
}

func (h *UserHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	type UserVerifyReq struct {
		Username string
		Password string
	}

	var req UserVerifyReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenstr, err := h.svc.VerifyUser(req.Username, req.Password)

	if err != nil {
		render(w, err, http.StatusInternalServerError)
	}

	w.Write([]byte(tokenstr))
	render(w, "Success", http.StatusCreated)
}
