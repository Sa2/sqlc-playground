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
	err := dbio.InitPgxDBConnPool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer dbio.GetPgxConnPool().Close()
	conn := db.New(dbio.GetPgxConnPool())
	users, err := conn.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
