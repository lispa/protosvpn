package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	authPackage "protosvpn-api/internal/auth"
)

func AdminMiddleware(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		tokenString := r.Header.Get(
			"Authorization",
		)

		if tokenString == "" {
			http.Error(
				w,
				"missing authorization header",
				http.StatusUnauthorized,
			)

			return
		}

		tokenString = tokenString[7:]

		token, err := jwt.Parse(
			tokenString,
			func(
				token *jwt.Token,
			) (interface{}, error) {
				return authPackage.GetJWTSecret(), nil
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

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(
				w,
				"invalid claims",
				http.StatusUnauthorized,
			)

			return
		}

		role, ok := claims["role"].(string)

		if !ok || role != "admin" {
			http.Error(
				w,
				"admin access required",
				http.StatusForbidden,
			)

			return
		}

		next(w, r)
	}
}
