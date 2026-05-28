package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type LoggerConfig struct {
	Level string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (LoggerConfig, error) {
	var config LoggerConfig

	if err := envconfig.Process("LOGGER", &config); err != nil {
		return LoggerConfig{}, fmt.Errorf("failed to process logger config: %w", err)
	}

	return config, nil
}	

func NewConfigMust() LoggerConfig {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("failed to load logger config: %w", err)
		panic(err)
	}
	return config
}