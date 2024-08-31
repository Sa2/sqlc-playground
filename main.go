package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	slog.Info("booting...")
	_, err := getDBConn()
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	PostgresUser string `envconfig:"POSTGRES_USER"`
	PostgresPass string `envconfig:"POSTGRES_PASSWORD"`
	PostgresHost string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort string `envconfig:"POSTGRES_PORT" default:"5432"`
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

func getDBConn() (*sql.DB, error) {
	c := getConfig()
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.PostgresUser, c.PostgresPass, c.PostgresHost, c.PostgresPort, c.PostgresDB)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	return conn, nil
}
