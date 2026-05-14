package main

import (
	"encoding/json"
	"log"
	"net/http"
	"protosvpn-api/internal/database"
	"protosvpn-api/internal/handlers"
	"protosvpn-api/internal/handlers/auth"
	"protosvpn-api/internal/middleware"

	adminHandlers "protosvpn-api/internal/handlers/admin"
	vpnHandlers "protosvpn-api/internal/handlers/vpn"
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
	http.HandleFunc("/api/v1/vpn/status", middleware.JWTAuthMiddleware(handlers.VPNStatusHandler))
	http.HandleFunc("/api/v1/auth/register", auth.RegisterHandler)
	http.HandleFunc("/api/v1/auth/login", auth.LoginHandler)
	http.HandleFunc("/api/v1/vpn/create-client", middleware.JWTAuthMiddleware(vpnHandlers.CreateClientHandler))
	http.HandleFunc("/api/v1/vpn/download-client", middleware.JWTAuthMiddleware(vpnHandlers.DownloadClientHandler))
	http.HandleFunc("/api/v1/vpn/revoke-client", middleware.AdminMiddleware(vpnHandlers.RevokeClientHandler))
	http.HandleFunc("/api/v1/vpn/clients", middleware.AdminMiddleware(vpnHandlers.ListClientsHandler))
	http.HandleFunc("/api/v1/admin/users", middleware.AdminMiddleware(adminHandlers.ListUsersHandler))
	http.HandleFunc("/api/v1/auth/me", middleware.JWTAuthMiddleware(auth.CurrentUserHandler))

	log.Println("API server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
