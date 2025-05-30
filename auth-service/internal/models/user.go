package models

// User type
type User struct {
	ID       string
	Email    string
	Password string
}

// UserStore interface to interact with the User DB
type UserStore interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}
