package vpn

import (
	"encoding/json"
	"log"
	"net/http"

	vpnService "protosvpn-api/internal/vpn"
)

type RevokeClientRequest struct {
	Name string `json:"name"`
}

func RevokeClientHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var request RevokeClientRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	err = vpnService.RevokeClient(
		request.Name,
	)

	if err != nil {
		log.Println(err)

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
		map[string]string{
			"message": "client revoked",
		},
	)
}
