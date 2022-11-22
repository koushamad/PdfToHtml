package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panicln(err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
