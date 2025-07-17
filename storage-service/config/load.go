package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(filenames ...string) error {
	return godotenv.Load(filenames...)
}

func GetVar(key string) string {
	return os.Getenv(key)
}