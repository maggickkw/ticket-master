package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maggickkw/ticket-master/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cofig *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(`host=%s user=%s dbname=%s password=%s sslmode=%s port=5432`,
		config.NewEnvConfig().DBHost, config.NewEnvConfig().DBUser, config.NewEnvConfig().DBName, config.NewEnvConfig().DBPassword, config.NewEnvConfig().DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to the database: %e", err)
	}

	log.Info("Connected to the database!")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate tables: %e", err)
	}

	return db
}
