package db

import (
	"context"
	"fmt"

	"github.com/Giovanni-Hernandez10/app-backend/auth-service/internal/models"
	"github.com/jackc/pgx/v5"
)

// has the reference to the database
type PostgresUserStore struct {
	DB *pgx.Conn
}

// implementing the CreateUser() function that is described in our UserStore interface
func (store *PostgresUserStore) CreateUser(ctx context.Context, user *models.User) error {
	// structuring the query to insert a new user to the table
	query := `INSERT INTO users (email, password) VALUES (@email, @password)`
	args := pgx.NamedArgs{
		"email":    user.Email,
		"password": user.Password,
	}

	// executing the query
	_, err := store.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("Unable to insert row: %w", err)
	}

	return nil
}

func (store *PostgresUserStore) GetUserByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	// structruing the query by email
	query := `SELECT id, email, password FROM users WHERE email = @email`
	args := pgx.NamedArgs{
		"email": user.Email,
	}

	// sending the query to the DB
	row, err1 := store.DB.Query(ctx, query, args)
	if err1 != nil {
		return nil, fmt.Errorf("unable to query user by email: %w", err1)
	}

	defer row.Close()

	// binding the data we received from query to the User struct
	user, err2 := pgx.CollectOneRow(row, pgx.RowToStructByName[*models.User])
	if err2 != nil {
		return nil, fmt.Errorf("unable to collect the user row: %w", err2)
	}

	return user, nil
}
