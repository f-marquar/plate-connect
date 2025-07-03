package handlers

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Kennzeichen string `json:"kennzeichen"`
	Email       string `json:"email"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// einfache Validierung
	if req.Kennzeichen == "" || req.Email == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
}
