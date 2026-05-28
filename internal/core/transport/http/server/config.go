package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)


type Config struct {
	Addr string `envconfig:"HTTP_ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to process HTTP server config: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load HTTP server config: %v", err))
	}

	return cfg
}
