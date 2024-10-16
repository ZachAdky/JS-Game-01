package main

import (
	"log"
	"net/http"

	"github.com/ZachAdky/JS-Game-01/backend/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/game", handlers.GameHandler).Methods("GET")

	log.Println("Server is running on port 8000...")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
