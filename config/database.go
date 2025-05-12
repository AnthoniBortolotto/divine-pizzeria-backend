package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConn() (*gorm.DB, error) {
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	user := GetEnv("DB_USERNAME")
	password := GetEnv("DB_PASSWORD")
	dbname := GetEnv("DB_NAME")
	sslmode := GetEnv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Error opening database connection")
		return nil, err
	}

	println("Successfully connected to database")
	return db, nil
}
