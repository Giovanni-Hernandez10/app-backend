package main

import (
	"context"
	"log"

	db "github.com/Giovanni-Hernandez10/app-backend/auth-service/internal/db"
	"github.com/joho/godotenv"
)

// entry point for the auth server
func main() {

	// loading env variables to use in our auth server
	godotenv.Load("../../../.env")

	// setting up db connection
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("There was an error connecting to the database: %v", err)
	}

	UserStore := &db.PostgresUserStore{
		DB: conn,
	}

	defer conn.Close(context.Background())
}
