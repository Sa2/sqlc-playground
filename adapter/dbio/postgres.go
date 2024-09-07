package dbio

import (
	"context"
	"fmt"

	"github.com/Sa2/sqlc-playground/adapter/env"
	"github.com/Sa2/sqlc-playground/gen/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitPgxDBConnPool(ctx context.Context) error {
	c := env.GetConfig()

	maxConn := 10
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d", c.PostgresUser, c.PostgresPass, c.PostgresHost, c.PostgresPort, c.PostgresDB, maxConn)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to open db connection: %w", err)
	}

	if pool == nil {
		return fmt.Errorf("pool is nil")
	}

	return nil
}
func GetPgxConnPool() *pgxpool.Pool {
	return pool
}

func GetSQLC() *db.Queries {
	return db.New(pool)
}
