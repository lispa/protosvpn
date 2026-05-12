package database

import "context"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func GetUsers() ([]User, error) {
	rows, err := DB.Query(
		context.Background(),
		`
		SELECT id, username, role
		FROM users
		ORDER BY id ASC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
