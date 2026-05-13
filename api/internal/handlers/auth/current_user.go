package auth

import (
	"encoding/json"
	"net/http"
	authPackage "protosvpn-api/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type CurrentUserResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func CurrentUserHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	tokenString :=
		r.Header.Get("Authorization")

	tokenString = tokenString[7:]

	token, err := jwt.Parse(
		tokenString,
		func(
			token *jwt.Token,
		) (interface{}, error) {
			return authPackage.GetJWTSecret(), nil
		},
	)

	if err != nil {
		http.Error(
			w,
			"invalid token",
			http.StatusUnauthorized,
		)

		return
	}

	claims :=
		token.Claims.(jwt.MapClaims)

	response := CurrentUserResponse{
		Username: claims["username"].(string),

		Role: claims["role"].(string),
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}
