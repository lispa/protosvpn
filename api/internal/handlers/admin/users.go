package admin

import (
	"encoding/json"
	"net/http"

	"protosvpn-api/internal/database"
)

func ListUsersHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	users, err := database.GetUsers()

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
			"users": users,
		},
	)
}
