package database

import "context"

func GetVPNClientOwner(
	clientName string,
) (int, error) {
	query := `
	SELECT owner_user_id
	FROM vpn_clients
	WHERE name = $1
	`

	var ownerUserID int

	err := DB.QueryRow(
		context.Background(),
		query,
		clientName,
	).Scan(&ownerUserID)

	if err != nil {
		return 0, err
	}

	return ownerUserID, nil
}
