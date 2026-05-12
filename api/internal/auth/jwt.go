package auth

import (
	"time"

	"protosvpn-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(
	username string,
	role string,
) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	jwtSecret := []byte(config.GetJWTSecret())

	return token.SignedString(jwtSecret)
}

func GetJWTSecret() []byte {
	return []byte(config.GetJWTSecret())
}
