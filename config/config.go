package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Host          string
	Port          string
	RedisHost     string
	RedisPort     string
	RedisPassword string
)

func LoadEnv() {
	log.Println("start loading env")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")

	log.Println("load env finished")
}
