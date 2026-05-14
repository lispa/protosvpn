package database

import "context"

type VPNClient struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func GetVPNClientsByUserID(
	userID int,
) ([]VPNClient, error) {
	query := `
	SELECT
		name,
		status
	FROM vpn_clients
	WHERE owner_user_id = $1
	ORDER BY id DESC
	`

	rows, err := DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []VPNClient

	for rows.Next() {
		var client VPNClient

		err := rows.Scan(
			&client.Name,
			&client.Status,
		)

		if err != nil {
			return nil, err
		}

		clients = append(
			clients,
			client,
		)
	}

	return clients, nil
}
