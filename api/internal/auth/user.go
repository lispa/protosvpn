package auth

import (
	"context"

	"protosvpn-api/internal/database"
)

func GetUserIDByUsername(
	username string,
) (int, error) {
	query := `
	SELECT id
	FROM users
	WHERE username = $1
	`

	var userID int

	err := database.DB.QueryRow(
		context.Background(),
		query,
		username,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}
