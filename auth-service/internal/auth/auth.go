package auth

import (
	"context"
	"strings"

	"github.com/Giovanni-Hernandez10/app-backend/auth-service/internal/db"
	pb "github.com/Giovanni-Hernandez10/app-backend/auth-service/proto/authpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// embedding the struct from proto code to create the grpc functions that it has associated with it
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	Store db.PostgresUserStore
}

// signup request logic
func (authServer *AuthServer) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.AuthResponse, error) {
	/* Steps Needed
	1. get all needed fields (email, password)
	2. validate input
		- ensure all fields exist
		- make sure passwords are the same
	3. check if it is already an existing user (check user DB)
	4. hash the password
	5. create a new user and store in DB
	6. return success
	*/

	user_email := req.GetEmail()
	user_password := req.GetPassword()
	confirm_password := req.GetConfirmPassword()

	// input validation checks
	if user_email == "" || user_password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email and password are required")
	}

	if len(user_password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "Password needs to be at least 8 characters long")
	}

	if user_password != confirm_password {
		return nil, status.Errorf(codes.InvalidArgument, "Passwords don't match")
	}

	if !strings.Contains(user_email, "@") {
		return nil, status.Errorf(codes.InvalidArgument, "Email is not formatted correctly")
	}

	// checking if the user already exists in our DB
	exists, err := authServer.Store.UserExists(ctx, user_email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error checking if user already exists: %v", err)
	}
	if exists {
		return nil, status.Errorf(codes.AlreadyExists, "User with email already exists")
	}

	// create the user by storing it in the DB
	err = authServer.Store.CreateUser(ctx, user_email, user_password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create the user: %v", err)
	}

	return &pb.AuthResponse{
		Success: true,
	}, nil
}

// // login request logic
// func (authServer *AuthServer) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {

// }

// // refresh request logic
// func (authServer *AuthServer) Refresh(context.Context, *pb.RefreshRequest) (*pb.RefreshResponse, error) {

// }

// // logout request logic
// func (authServer *AuthServer) Logout(context.Context, *pb.LogoutRequest) (*pb.AuthResponse, error) {

// }

// // forgot password request logic
// func (authServer *AuthServer) ForgotPassword(context.Context, *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {

// }
