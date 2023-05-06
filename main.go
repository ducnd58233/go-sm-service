package main

import (
	"context"
	"fmt"
	"go-sm-service/component"
	"go-sm-service/config"
	"go-sm-service/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger.Info("Starting up...")
	c, err := config.New()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	// use the configuration values here
	dbUser := c.DB.User
	dbPass := c.DB.Password
	dbHost := c.DB.Host
	dbPort := c.DB.Port
	dbName := c.DB.DBName
	rdAddr := c.Redis.Addr
	rdPass := c.Redis.Password
	rdDB := c.Redis.DB

	// Connect to MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, errCon := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errCon != nil {
		logger.Error(errCon)
		panic(err)
	}

	logger.Info("Connected to database successful.")

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdAddr,
		Password: rdPass,
		DB:       rdDB,
	})
	ctx := context.Background() // Create a new context
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		logger.Error(errCon)
		panic(err)
	}

	logger.Info("Connected to Redis successful.")

	router := gin.Default()

	appCtx := component.NewAppContext(db)

	logger.Info(appCtx)
	router.Run(":6000")
	logger.Info("Server started.")
}
