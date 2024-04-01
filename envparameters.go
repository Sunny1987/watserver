package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetEnvValueFor function will get all necessary variables for the app
func GetEnvValueFor(key string) string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
