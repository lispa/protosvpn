package database

import (
	"context"
	"log"
)

func RunMigrations() {
	usersQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'user',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(
		context.Background(),
		usersQuery,
	)

	if err != nil {
		log.Fatal("Failed to run users migration:", err)
	}

	vpnClientsQuery := `
	CREATE TABLE IF NOT EXISTS vpn_clients (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		owner_user_id INTEGER REFERENCES users(id),
		status TEXT NOT NULL DEFAULT 'active',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		revoked_at TIMESTAMP
	);
	`

	_, err = DB.Exec(
		context.Background(),
		vpnClientsQuery,
	)

	if err != nil {
		log.Fatal(
			"Failed to run vpn clients migration:",
			err,
		)
	}

	log.Println("Database migrations completed")
}
