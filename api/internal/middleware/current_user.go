package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	authPackage "protosvpn-api/internal/auth"
)

type CurrentUser struct {
	Username string
	Role     string
}

func GetCurrentUser(
	r *http.Request,
) (*CurrentUser, error) {
	authHeader := r.Header.Get(
		"Authorization",
	)

	if authHeader == "" {
		return nil, errors.New(
			"missing authorization header",
		)
	}

	if !strings.HasPrefix(
		authHeader,
		"Bearer ",
	) {
		return nil, errors.New(
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
		return nil, err
	}

	claims, ok :=
		token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New(
			"invalid token claims",
		)
	}

	username, ok :=
		claims["username"].(string)

	if !ok {
		return nil, errors.New(
			"username missing",
		)
	}

	role, ok :=
		claims["role"].(string)

	if !ok {
		return nil, errors.New(
			"role missing",
		)
	}

	return &CurrentUser{
		Username: username,
		Role:     role,
	}, nil
}
