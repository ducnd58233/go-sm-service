package config

import (
	"errors"
	"go-sm-service/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Redis struct {
	Addr     string
	Password string
	DB       int
}

func (c *Redis) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	var ok bool

	// Get the Redis configuration constants
	c.Addr, ok = os.LookupEnv("REDIS_ADDR")
	if !ok {
		logger.Error("Missing Redis Address string.")
		return errors.New("missing environment variable(s)")
	}

	c.Password, ok = os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		logger.Error("Missing Redis Password string.")
		logger.Info("Using default password: ''.")
		c.Password = ""
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		c.DB = 0 // set default value
	} else {
		c.DB = db
	}

	return nil
}
