package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Data represents the structure of the JSON response
type Data struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	Message   string  `json:"message"`
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// Create a new data object
		data := Data{
			Timestamp: time.Now().Format(time.RFC3339),
			Value:     rand.Float64() * 100, // Generate a random value
			Message:   "Real-time data from Go backend!",
		}

		// Set the response header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Encode the data object into JSON and send it
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		}
	})

	// Start the server
	port := "8080"
	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
