package database

import "context"

func CreateVPNClient(
	name string,
	ownerUserID int,
) error {
	query := `
	INSERT INTO vpn_clients (
		name,
		owner_user_id,
		status
	)
	VALUES ($1, $2, 'active')
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		name,
		ownerUserID,
	)

	return err
}
