package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// has the reference to the database
type PostgresUserStore struct {
	DB *pgx.Conn
}

// implementing the CreateUser() function that is described in our UserStore interface
func (store *PostgresUserStore) CreateUser(ctx context.Context, email string, password string) error {
	// structuring the query to insert a new user to the table
	query := `INSERT INTO "User" (email, password) VALUES (@email, @password)`
	args := pgx.NamedArgs{
		"email":    email,
		"password": password,
	}

	// executing the query
	_, err := store.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

// checking to see if the user already exists in our table
func (store *PostgresUserStore) UserExists(ctx context.Context, email string) (bool, error) {
	// structuring the query by email
	query := `SELECT 1 FROM "User" WHERE email = @email`
	args := pgx.NamedArgs{
		"email": email,
	}

	row := store.DB.QueryRow(ctx, query, args)

	// checking if the user already exists in our table
	var exists int
	err := row.Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("error checking if user already exists: %w", err)
	}

	return true, nil
}

// func (store *PostgresUserStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
// 	// structruing the query by email
// 	query := `SELECT id, email, password FROM "User" WHERE email = @email`
// 	args := pgx.NamedArgs{
// 		"email": email,
// 	}

// 	// sending the query to the DB
// 	row, err1 := store.DB.Query(ctx, query, args)
// 	if err1 != nil {
// 		return nil, fmt.Errorf("unable to query user by email: %w", err1)
// 	}

// 	defer row.Close()

// 	// binding the data we received from query to the User struct
// 	user, err2 := pgx.CollectOneRow(row, pgx.RowToStructByName[models.User])
// 	if err2 != nil {
// 		return nil, fmt.Errorf("unable to collect the user row: %w", err2)
// 	}

// 	return &user, nil
// }
