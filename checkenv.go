package main

import (
	"github.com/joho/godotenv"
	"github.com/thep0y/go-logger/log"
	"os"
)

func checkDotEnv() bool {
	if err := godotenv.Load(os.ExpandEnv("$HOME/.portmonitor/.env")); err != nil {
		log.Fatal("No .env file found")
	}

	return false
}
