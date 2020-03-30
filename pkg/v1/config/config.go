package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER" default:"postgres"`
	DBName   string `envconfig:"DB_NAME" default:"amazing-payments"`
	DBPassword string `envconfig:"DB_PASS" default:"postgres"`
}

func InitializeConfig() *Config {
	var c Config
	err := envconfig.Process("amazing-payments", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &c
}
