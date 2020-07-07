package main

import (
	"encoding/json"
	"net/http"
)

type request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func health(w http.ResponseWriter, r *http.Request) {
	var response = map[string]string{
		"message": "pong",
	}
	respondWithJSON(w, http.StatusOK, "SUCCESS", response)
}

func add(w http.ResponseWriter, r *http.Request) {
	var req request
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "BAD_REQUEST", nil)
		return
	}
	respondWithJSON(w, http.StatusOK, "SUCCESS", req)
}
