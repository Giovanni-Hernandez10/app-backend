package main

import (
	"context"
	"log"
	"net"

	"github.com/Giovanni-Hernandez10/app-backend/auth-service/internal/auth"
	db "github.com/Giovanni-Hernandez10/app-backend/auth-service/internal/db"
	pb "github.com/Giovanni-Hernandez10/app-backend/auth-service/proto/authpb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

	defer conn.Close(context.Background())

	// setting up tcp connection and grpc server
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("There was an error listening: %v", err)
	}
	grpcServer := grpc.NewServer()

	store := &db.PostgresUserStore{
		DB: conn,
	}

	pb.RegisterAuthServiceServer(grpcServer, &auth.AuthServer{
		Store: *store,
	})

	// this is what wiill receive the incoming requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
