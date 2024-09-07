package main

import (
	"context"
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
	uid := ulid.Make().String()
	_, err = txQuery.CreateUser(ctx, db.CreateUserParams{
		ID:       uid,
		Name:     "test",
		Email:    fmt.Sprintf("%s@domain", ulid.Make().String()),
		Password: "password",
	})
	if err != nil {
		fmt.Errorf("failed to create user: %w", err)
	}
	_, err = txQuery.CreateUserDetail(ctx, db.CreateUserDetailParams{
		ID:          ulid.Make().String(),
		UserID:      uid,
		DetailInfo1: "detail1",
		DetailInfo2: "detail2",
	})
	if err != nil {
		fmt.Errorf("failed to create user detail: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	fmt.Println("transaction committed")

	return nil
}
