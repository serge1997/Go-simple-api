package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_NAME = ""
	APP_PORT = ""
)

func LoadEnv() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	APP_PORT = os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "4800"
	}
	APP_NAME = os.Getenv("APP_NAME")
}
