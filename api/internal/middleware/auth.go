package middleware

import (
	"net/http"
	"strings"

	"protosvpn-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(
				w,
				"missing authorization header",
				http.StatusUnauthorized,
			)

			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				jwtSecret := []byte(config.GetJWTSecret())

				return jwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(
				w,
				"invalid token",
				http.StatusUnauthorized,
			)

			return
		}

		next(w, r)
	}
}
