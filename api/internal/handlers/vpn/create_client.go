package vpn

import (
	"encoding/json"
	"net/http"

	vpnService "protosvpn-api/internal/vpn"
)

type CreateClientRequest struct {
	Name string `json:"name"`
}

type CreateClientResponse struct {
	Message string `json:"message"`
}

func CreateClientHandler(
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

	var request CreateClientRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	err = vpnService.CreateClient(
		request.Name,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	response := CreateClientResponse{
		Message: "client created",
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}
