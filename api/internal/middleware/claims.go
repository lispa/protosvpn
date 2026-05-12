package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	authPackage "protosvpn-api/internal/auth"
)

func GetUsernameFromRequest(
	r *http.Request,
) (string, error) {
	authHeader := r.Header.Get(
		"Authorization",
	)

	if authHeader == "" {
		return "", errors.New(
			"missing authorization header",
		)
	}

	if !strings.HasPrefix(
		authHeader,
		"Bearer ",
	) {
		return "", errors.New(
			"invalid authorization header",
		)
	}

	tokenString := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	token, err := jwt.Parse(
		tokenString,
		func(
			token *jwt.Token,
		) (interface{}, error) {
			return authPackage.GetJWTSecret(), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok :=
		token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New(
			"invalid token claims",
		)
	}

	username, ok :=
		claims["username"].(string)

	if !ok {
		return "", errors.New(
			"username not found in token",
		)
	}

	return username, nil
}
