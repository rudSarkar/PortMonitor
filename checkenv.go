package main

import (
	"github.com/joho/godotenv"
	"github.com/thep0y/go-logger/log"
)

func checkDotEnv() bool {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}

	return false
}
