package postgres

import (
	"context"
	"fmt"

	"avito-internship/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewClient ...
func NewClient(ctx context.Context, config *config.Postgres) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName)
	conn, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("can't ping database: %w", err)
	}
	return conn, nil
}
