package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the serverxxxxxxxxxxx!")
	})

	// Log that the server is starting
	log.Println("Starting server on port 8080...")

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
