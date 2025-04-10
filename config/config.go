package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system ENV.")
	}
}

func EnvPort() string {
	return os.Getenv("PORT")
}

func EnvDBConnection() string {
	return os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func SetupLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
}
