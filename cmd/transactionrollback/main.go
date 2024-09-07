package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"

	"github.com/Sa2/sqlc-playground/adapter/dbio"
	"github.com/Sa2/sqlc-playground/gen/db"
	"github.com/oklog/ulid/v2"
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
	query := db.New(pool)
	users, err := query.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)
	txQuery := query.WithTx(tx)
	_, err = txQuery.CreateUser(ctx, db.CreateUserParams{
		ID:       string(ulid.Make().String()),
		Name:     "test",
		Email:    fmt.Sprintf("%s@domain", ulid.Make().String()),
		Password: "password",
	})
	if err != nil {
		fmt.Errorf("failed to create user: %w", err)
	}
	return errors.New("error occurred")
}
