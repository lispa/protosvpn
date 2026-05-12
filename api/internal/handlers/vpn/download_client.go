package vpn

import (
	"net/http"

	vpnService "protosvpn-api/internal/vpn"
)

func DownloadClientHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	clientName := r.URL.Query().Get("name")

	if clientName == "" {
		http.Error(
			w,
			"missing client name",
			http.StatusBadRequest,
		)

		return
	}

	config, err := vpnService.GetClientConfig(
		clientName,
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
		"application/octet-stream",
	)

	w.Header().Set(
		"Content-Disposition",
		"attachment; filename=\""+clientName+".ovpn\"",
	)

	w.Write(config)
}
