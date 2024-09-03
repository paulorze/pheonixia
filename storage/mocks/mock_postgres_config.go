package mocks

import (
	"log"
	"os"
	"phoenixia/storage"

	"github.com/joho/godotenv"
)

func MockPostgresConnect() *storage.Config {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		log.Fatal(err)
	}
	config := storage.Config{
		Host:     os.Getenv("DB_TEST_HOST"),
		Port:     os.Getenv("DB_TEST_PORT"),
		Password: os.Getenv("DB_TEST_PASS"),
		User:     os.Getenv("DB_TEST_USER"),
		SSLMode:  os.Getenv("DB_TEST_SSLMODE"),
		DBName:   os.Getenv("DB_TEST_NAME"),
	}

	return &config
}
