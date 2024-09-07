package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresUser string `envconfig:"POSTGRES_USER"`
	PostgresPass string `envconfig:"POSTGRES_PASSWORD"`
	PostgresHost string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort uint16 `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDB   string `envconfig:"POSTGRES_DB"`
}

func GetConfig() Config {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
