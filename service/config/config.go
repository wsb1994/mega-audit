package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	PG_HOST     string `json:"pg_host"`
	PG_PORT     int    `json:"pg_port"`
	PG_USER     string `json:"pg_user"`
	PG_PASSWORD string `json:"pg_password"`
	PG_DATABASE string `json:"pg_database"`
	APP_PORT    string `json:"app_port"`
}

func Loadconfig() *PostgresConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		os.Exit(1)
	}
	return &PostgresConfig{
		PG_HOST:     os.Getenv("PG_HOST"),
		PG_PORT:     port,
		PG_USER:     os.Getenv("PG_USER"),
		PG_PASSWORD: os.Getenv("PG_PASSWORD"),
		PG_DATABASE: os.Getenv("PG_DATABASE"),
		APP_PORT:    os.Getenv("APP_PORT"),
	}
}
