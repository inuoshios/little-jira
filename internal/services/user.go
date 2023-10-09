package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/inuoshios/little-jira/internal/models"
)

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}

func CreateUser(user *models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users
	 (username, first_name, last_name, email, password, gender)
	 values ($1, $2, $3, $4, $5, $6)
	 returning id
	`

	err := db.QueryRowContext(ctx, query,
		user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.Gender).Scan(&user.ID)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}

	return user.ID, nil
}

func GetUser(payload *models.Login) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	query := `
		select id, first_name, last_name, password, username, email, gender, created_at, updated_at 
		from users where email = $1 or username = $2;`
	rows := db.QueryRowContext(ctx, query, payload.Email, payload.Username)
	if err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Username,
		&user.Email,
		&user.Gender,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, first_name, last_name, username, email, gender, created_at, updated_at from users;`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.Gender,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
