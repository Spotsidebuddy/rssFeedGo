package main

import (
	"net/http"
)

func handlerErr(writer http.ResponseWriter, response *http.Request) {
	RespondWithError(writer, 400, "Something went wrong")
}