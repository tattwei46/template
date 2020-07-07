package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	HttpStatus int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func respondWithJSON(w http.ResponseWriter, code int, message string, data interface{}) {

	r := Response{
		HttpStatus: code,
		Message:    message,
		Data:       data,
	}
	response, _ := json.Marshal(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
