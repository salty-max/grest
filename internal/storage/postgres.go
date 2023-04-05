package storage

import (
	"fmt"

	"github.com/salty-max/grest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func Connect(env config.EnvVars) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Europe/Paris",
		env.DB_HOST,
		env.DB_PORT,
		env.DB_NAME,
		env.DB_USER,
		env.DB_PASSWORD,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate()

	return &Database{
		Client: DB,
	}, nil
}

func Close(db *Database) error {
	sqlDB, err := db.Client.DB()
	if err != nil {
		return err
	}

	sqlDB.Close()
	return nil
}
