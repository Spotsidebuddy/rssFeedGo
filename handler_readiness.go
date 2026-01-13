package main

import (
	"net/http"
)

func handlerReadiness(writer http.ResponseWriter, response *http.Request) {
	respondWithJSON(writer, 200, struct{Server string `json:"server"`}{Server: "Ready"})
}