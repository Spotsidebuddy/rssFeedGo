package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall JSON response: %v", payload)
		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data) 
}

func RespondWithError(writer http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with error 5**: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(writer, code, errResponse{Error: msg})
}