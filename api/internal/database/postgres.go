package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	databaseURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = conn

	log.Println("Connected to PostgreSQL")
}
