package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load values from .env into the system.
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value.
	// Outputs a panic if the value is missing.
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panic("Value " + v + "does not exist")
	}
	return value
}
