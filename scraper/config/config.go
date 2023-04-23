package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

type envConfig map[string]string

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Config struct {
	Database DBConfig
	Port     int
}

// New returns a new Config struct
func New() *Config {
	var (
		err       error
		envConfig envConfig
	)

	envConfig = initEnvConfig()
	err = configCheck(envConfig)

	if err != nil {
		err = godotenv.Load("env/.env")
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		envConfig = initEnvConfig()
		err = configCheck(envConfig)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}

	return &Config{
		Database: DBConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			Username: getEnv("DB_USERNAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_DATABASE", "postgres"),
		},
		Port: getEnvAsInt("PORT", 8080),
	}
}

func initEnvConfig() envConfig {
	envConfig := make(envConfig)

	envConfig["Port"] = os.Getenv("PORT")

	envConfig["DBHost"] = os.Getenv("DB_HOST")
	envConfig["DBPort"] = os.Getenv("DB_PORT")
	envConfig["DBUsername"] = os.Getenv("DB_USERNAME")
	envConfig["DBPassword"] = os.Getenv("DB_PASSWORD")
	envConfig["DBDatabase"] = os.Getenv("DB_DATABASE")

	return envConfig
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func configCheck(envConfig envConfig) error {
	isEnvNotFull := false
	errorText := "Set full list of default OS variables or create 'env/.env' file with it. Empty variables:\n"

	for key, value := range envConfig {
		if value == "" {
			isEnvNotFull = true
			errorText = errorText + key + "\n"
		}
	}

	if isEnvNotFull {
		return errors.New(errorText)
	}
	return nil
}
