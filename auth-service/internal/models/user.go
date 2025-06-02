package models

import "context"

// User type
type User struct {
	ID       int8
	Email    string
	Password string
}

// UserStore interface to interact with the User DB
type UserStore interface {
	CreateUser(ctx context.Context, email string, password string) error
	UserExists(ctx context.Context, email string) (bool, error)
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
}
