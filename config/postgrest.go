package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{}

func (c *Config) DB() *gorm.DB {
	Env := godotenv.Load()

	if Env != nil {
		log.Panic("Failed to load env file!")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port 5432 sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect database")
	}

	return db
}

func NewDB() *Config {
	return &Config{}
}
