package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
)

type Config struct {
	DBPort     string
	DbUser     string
	DBPassword string
	DBHost     string
	DBName     string
	AppEnv     string
	Port       string
}

func Load(env string) {
	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
}

func Env() Config {
	return Config{
		DBPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		Port:       os.Getenv("PORT"),
	}
}

func GetDSN() string {
	dbUser := Env().DbUser
	dbPassword := url.QueryEscape(Env().DBPassword)
	dbName := Env().DBName
	dbPort := Env().DBPort
	dbHost := Env().DBHost
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}
