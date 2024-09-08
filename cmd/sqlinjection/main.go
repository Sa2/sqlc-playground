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
	injectionQuery := "01J779DFBWF5Q879XK73SH2FA2; DROP TABLE users;"
	user, err := conn.GetUserByID(ctx, injectionQuery)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	fmt.Println(user)
	return nil
}
