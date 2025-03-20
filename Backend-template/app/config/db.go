package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Nxwbtk/Mono-repo-template/Backend-template/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateClientDatabase() (*gorm.DB, *sql.DB, error) {
	config := NewConfig()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.POSTGRES_HOST,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_DB,
		config.POSTGRES_SSL,
		config.POSTGRES_TIMEZONE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error while creating connection to the database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("Error while getting the database connection!", err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(10)
	sqlDB.SetMaxOpenConns(10)

	err = sqlDB.Ping()
	if err != nil {
		log.Panic("Could not ping datanase")
	}

	return db, sqlDB, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Cat{})
	if err != nil {
		return fmt.Errorf("error while migrating the database: %w", err)
	}
	return nil
}
