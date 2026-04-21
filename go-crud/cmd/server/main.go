package main

import (
	"go-crud/internal/modules/user"
	"log"
	"net/http"
)

func main() {
	// Initialize repository, service, and handler
	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	// Set up HTTP routes and start the server
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handler.GetAll)
	mux.HandleFunc("POST /users", handler.Create)
	// Additional routes for GetByID, Update, Delete would go here

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
