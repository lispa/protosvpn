package vpn

import (
	"encoding/json"
	"net/http"

	authPackage "protosvpn-api/internal/auth"
	"protosvpn-api/internal/database"
	"protosvpn-api/internal/middleware"

	vpnService "protosvpn-api/internal/vpn"
)

func ListClientsHandler(
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

	if currentUser.Role == "admin" {
		clients, err :=
			vpnService.ListClients()

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"clients": clients,
			},
		)

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

	clients, err :=
		database.GetVPNClientsByUserID(
			userID,
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"clients": clients,
		},
	)
}
