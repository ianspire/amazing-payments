package pkg

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER" default:"postgres"`
	DBName     string `envconfig:"DB_NAME" default:"amazing-payments"`
	DBPassword string `envconfig:"DB_PASS" default:"postgres"`
	JSONPort   string `envconfig:"JSON_PORT" default:":8080"`
	RPCPort    string `envconfig:"RPC_PORT" default:":8081"`
	StripeKey  string `envconfig:"STRIPE_KEY" required:"true"`
}

func NewConfig() *Config {
	var c Config
	err := envconfig.Process("amazing-payments", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &c
}
