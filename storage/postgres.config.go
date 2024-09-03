package storage

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PostgresConnect() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return &config
}