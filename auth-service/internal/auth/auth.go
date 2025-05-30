package auth

import (
	"context"

	pb "github.com/Giovanni-Hernandez10/app-backend/auth-service/proto/authpb"
)

// embedding the struct from proto code to create the grpc functions that it has associated with it
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

// signup request logic
func (authServer *AuthServer) Signup(context.Context, *pb.SignupRequest) (*pb.AuthResponse, error) {
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
}

// login request logic
func (authServer *AuthServer) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {

}

// refresh request logic
func (authServer *AuthServer) Refresh(context.Context, *pb.RefreshRequest) (*pb.RefreshResponse, error) {

}

// logout request logic
func (authServer *AuthServer) Logout(context.Context, *pb.LogoutRequest) (*pb.AuthResponse, error) {

}

// forgot password request logic
func (authServer *AuthServer) ForgotPassword(context.Context, *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {

}
