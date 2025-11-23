package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPathKey = "--env-path"
)

type Postgres struct {
	DBName   string `envconfig:"DB_NAME" default:"pr-reviewer"`
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Username string `envconfig:"DB_USERNAME" default:"test"`
	Password string `envconfig:"DB_PASSWORD" default:"test"`
}

var (
	errEnvPathNotSet           = errors.New("--env-path environment variable not set")
	errInvalidProgramArguments = errors.New("invalid program arguments")
)

// FromEnv ...
func FromEnv() (*Postgres, error) {
	envPath, err := parseEnvPath()
	if err != nil {
		return nil, err
	}
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("can't load .env file: %w", err)
	}
	config := &Postgres{}
	if err = envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

func parseEnvPath() (string, error) {
	// Первая часть - имя программы
	if len(os.Args[1:]) != 2 {
		return "", errInvalidProgramArguments
	}
	if os.Args[1] != envPathKey {
		return "", errEnvPathNotSet
	}
	return os.Args[2], nil
}
