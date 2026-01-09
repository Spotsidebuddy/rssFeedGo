package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("No PORT found in environment variables")
	}
	fmt.Println("Port:", portString)

	router := chi.NewRouter()

	router.Use(
		cors.Handler(
			//someArgs Here
		)
	)

	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

