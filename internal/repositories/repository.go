package repositories

import (
	"log/slog"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db      *sqlx.DB
	logger  *slog.Logger
	redisCl *redis.Client
}

func NewRepository(db *sqlx.DB, redisCl *redis.Client, logger *slog.Logger) *Repository {
	return &Repository{db: db, logger: logger, redisCl: redisCl}
}
