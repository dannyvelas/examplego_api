package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
)

type Config struct {
	http     HttpConfig
	postgres PostgresConfig
	token    TokenConfig
}

const projectName = "examplego_api"

func loadDotEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`^(.*` + projectName + `)`)
	rootPath := re.Find([]byte(cwd))

	return godotenv.Load(string(rootPath) + `/.env`)
}

func NewConfig() Config {
	err := loadDotEnv()
	if err != nil {
		log.Warn().Msgf("config: .env file not found: %v", err)
	}

	return Config{
		http:     newHttpConfig(),
		postgres: newPostgresConfig(),
		token:    newTokenConfig(),
	}
}

func (config Config) Http() HttpConfig {
	return config.http
}

func (config Config) Postgres() PostgresConfig {
	return config.postgres
}

func (config Config) Token() TokenConfig {
	return config.token
}
