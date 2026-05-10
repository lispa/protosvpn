package main

import (
	"encoding/json"
	"log"
	"net/http"
	"protosvpn-api/internal/database"
	"protosvpn-api/internal/handlers"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	database.Connect()
	database.RunMigrations()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/v1/vpn/status", handlers.VPNStatusHandler)

	log.Println("API server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
