package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/Sa2/sqlc-playground/gen/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	ctx := context.Background()
	slog.Info("booting...")
	pool, err := newPool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	conn := db.New(pool)
	users, err := conn.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

type Config struct {
	PostgresUser string `envconfig:"POSTGRES_USER"`
	PostgresPass string `envconfig:"POSTGRES_PASSWORD"`
	PostgresHost string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort uint16 `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDB   string `envconfig:"POSTGRES_DB"`
}

func getConfig() Config {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func newPool(ctx context.Context) (*pgxpool.Pool, error) {
	c := getConfig()

	maxConn := 10
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d", c.PostgresUser, c.PostgresPass, c.PostgresHost, c.PostgresPort, c.PostgresDB, maxConn)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	return pool, nil
}
