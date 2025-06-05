package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	if env != "test" {
		godotenv.Load(".env.local") //nolint
	}
	godotenv.Load(".env." + env) //nolint
	godotenv.Load()              //nolint
}
