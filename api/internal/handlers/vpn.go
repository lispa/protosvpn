package handlers

import (
	"encoding/json"
	"net/http"

	"protosvpn-api/internal/openvpn"
)

func VPNStatusHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := openvpn.ParseStatusFile("/vpn-data/openvpn-status.log")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"clients": clients,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
