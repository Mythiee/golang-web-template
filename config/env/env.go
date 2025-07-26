package env

import (
	"os"

	"github.com/joho/godotenv"
)

// InitEnv If the variable named "APP_ENV" cannot be read or its value is not equal to "production",
// then read the environment variables from the ".env" file
func InitEnv() {
	if env := os.Getenv("APP_ENV"); env != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			// TODO log
			return
		}
	}
}
