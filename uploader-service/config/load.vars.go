package config

import (
	"os"

	"github.com/joho/godotenv"
)

// InitConfig will initialize the env configuation
// This will be mounted at the main application which acts like root point.
func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

// GetVaribale is used to retrieve the configuration varibales
// to be used cross in our application.
func GetVariable(key string) string {
	return os.Getenv(key)
}