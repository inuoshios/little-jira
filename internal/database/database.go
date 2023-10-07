package database

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectToDatabase(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.SetConnMaxLifetime(30 * time.Minute)
	conn.SetMaxOpenConns(15)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxIdleTime(10 * time.Minute)

	return conn, nil
}
