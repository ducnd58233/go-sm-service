package config

import (
	"errors"
	"go-sm-service/logger"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c *DB) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	var ok bool

	// Get the database configuration constants
	c.User, ok = os.LookupEnv("DB_USER")
	if !ok {
		logger.Error("Missing DB User string.")
		logger.Info("Using default user: root.")
		c.User = "root"
	}

	c.Password, ok = os.LookupEnv("DB_PASSWORD")
	if !ok {
		logger.Error("Missing DB Password string.")
		logger.Info("Using default password: ''.")
		c.Password = ""
	}

	c.Host, ok = os.LookupEnv("DB_HOST")
	if !ok {
		logger.Error("Missing DB Host string.")
		return errors.New("missing environment variable(s)")
	}

	c.Port, ok = os.LookupEnv("DB_PORT")
	if !ok {
		logger.Error("Missing DB Port string.")
		return errors.New("missing environment variable(s)")
	}

	c.DBName, ok = os.LookupEnv("DB_DBName")
	if !ok {
		logger.Error("Missing DB Database name string.")
		return errors.New("missing environment variable(s)")
	}

	return nil
}
