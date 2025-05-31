package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// connect to postgresql through supabase - returns a connection and error
func Connect() (*pgx.Conn, error) {
	db_url := os.Getenv("DATABASE_URL")

	// connecting to the database
	conn, err := pgx.Connect(context.Background(), db_url)

	if err != nil {
		log.Fatalf("Failed to connect to the database with error: %v", err)
		return nil, err
	}

	// checking to see if we are connected
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
		conn.Close(context.Background())
		return nil, err
	}

	log.Println("Connected to:", version)
	return conn, nil
}
