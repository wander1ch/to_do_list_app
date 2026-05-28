package core_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `envconfig:"HOST" required:"true"`
	Port int	`envconfig:"PORT" default:"5432"`
	User string	`envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DBName string `envconfig:"DB" required:"true"`
	Timeout time.Duration 	`envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config
	
	if err := envconfig.Process("POSTGRES", &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to process Postgres config: %w", err)
	}
	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("failed to load Postgres config: %w", err)
		panic(err)
	}
	return cfg
}