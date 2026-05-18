package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	internalAuth "protosvpn-api/internal/auth"
	"protosvpn-api/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(
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

	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	log.Println(request.Username)
	log.Println(request.Password)

	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	query := `
	SELECT password_hash, role
	FROM users
	WHERE username = $1
	`

	var passwordHash string
	var role string

	err = database.DB.QueryRow(
		context.Background(),
		query,
		request.Username,
	).Scan(
		&passwordHash,
		&role,
	)

	if err != nil {
		http.Error(
			w,
			"invalid credentials",
			http.StatusUnauthorized,
		)

		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(request.Password),
	)

	if err != nil {
		http.Error(
			w,
			"invalid credentials",
			http.StatusUnauthorized,
		)

		return
	}

	token, err := internalAuth.GenerateJWT(
		request.Username,
		role,
	)

	if err != nil {
		http.Error(
			w,
			"failed to generate token",
			http.StatusInternalServerError,
		)

		return
	}

	response := LoginResponse{
		Token: token,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}
