package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"protosvpn-api/internal/database"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func RegisterHandler(
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

	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(
			w,
			"failed to hash password",
			http.StatusInternalServerError,
		)

		return
	}

	query := `
	INSERT INTO users (
		username,
		password_hash
	)
	VALUES ($1, $2)
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		request.Username,
		string(passwordHash),
	)

	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to create user",
			http.StatusInternalServerError,
		)

		return
	}

	response := RegisterResponse{
		Message: "user created",
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}
