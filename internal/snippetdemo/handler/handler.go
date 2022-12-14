package handler

import (
	"encoding/json"
	"net/http"
)

func render(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	//** CORS SPECIFIC SECTION **//
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	//** **//

	w.WriteHeader(status)

	switch v := body.(type) {
	case string:
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: v,
		})
	case error:
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: v.Error(),
		})
	case nil:
		// do nothing
	default:
		json.NewEncoder(w).Encode(body)
	}
}
