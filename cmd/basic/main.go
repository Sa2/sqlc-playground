package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/Sa2/sqlc-playground/adapter/dbio"
	"github.com/Sa2/sqlc-playground/gen/db"
)

func main() {
	ctx := context.Background()
	slog.Info("booting...")
	err := app(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func app(ctx context.Context) error {
	err := dbio.InitPgxDBConnPool(ctx)
	if err != nil {
		return fmt.Errorf("failed to init pgx db conn pool: %w", err)
	}
	pool := dbio.GetPgxConnPool()
	defer pool.Close()
	conn := db.New(pool)
	users, err := conn.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
	return nil
}
