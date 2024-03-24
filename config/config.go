package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host   string
	Port   string
	DBUser string
	DBPass string
	DBAddr string
	DBName string
}

var Env = initConfig()

func initConfig() Config {

	// load environment variables
	godotenv.Load()

	return Config{
		Host:   getEnv("HOST", "http://localhost"),
		Port:   getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "dbeaver"),
		DBPass: getEnv("DB_PASS", "dbeaver"),
		DBAddr: getEnv("DB_ADDR", "127.0.0.1:3306"),
		DBName: getEnv("DB_NAME", "go-ecom"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
