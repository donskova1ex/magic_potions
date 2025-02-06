package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(pgDSN string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", pgDSN)

	if err != nil {
		return nil, fmt.Errorf("failed to open connection to postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return db, nil
}
