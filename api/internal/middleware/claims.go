package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	authPackage "protosvpn-api/internal/auth"
)

func GetUsernameFromRequest(
	r *http.Request,
) (string, error) {
	tokenString := r.Header.Get(
		"Authorization",
	)

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
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	username :=
		claims["username"].(string)

	return username, nil
}
