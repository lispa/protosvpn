package handlers

import (
	"encoding/json"
	"net/http"

	authPackage "protosvpn-api/internal/auth"
	"protosvpn-api/internal/database"
	"protosvpn-api/internal/middleware"
	"protosvpn-api/internal/openvpn"
)

func VPNStatusHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	currentUser, err :=
		middleware.GetCurrentUser(r)

	if err != nil {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)

		return
	}

	clients, err :=
		openvpn.ParseStatusFile(
			"/vpn-data/openvpn-status.log",
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	if currentUser.Role == "admin" {
		response := map[string]interface{}{
			"clients": clients,
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		json.NewEncoder(w).Encode(response)

		return
	}

	userID, err :=
		authPackage.GetUserIDByUsername(
			currentUser.Username,
		)

	if err != nil {
		http.Error(
			w,
			"failed to get user",
			http.StatusInternalServerError,
		)

		return
	}

	var filteredClients []openvpn.Client

	for _, client := range clients {
		ownerUserID, err :=
			database.GetVPNClientOwner(
				client.Name,
			)

		if err != nil {
			continue
		}

		if ownerUserID == userID {
			filteredClients = append(
				filteredClients,
				client,
			)
		}
	}

	response := map[string]interface{}{
		"clients": filteredClients,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}
