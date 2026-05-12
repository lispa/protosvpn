package vpn

import (
	"encoding/json"
	"net/http"

	vpnService "protosvpn-api/internal/vpn"
)

func ListClientsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	clients, err := vpnService.ListClients()

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
