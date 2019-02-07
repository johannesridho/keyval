package redisprovider

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/johannesridho/keyval/config"
	"log"
	"os"
)

var client *redis.Client

func init() {
	log.Println("initiating Redis")

	config.LoadEnv()

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	address := fmt.Sprintf("%s:%s", host, port)

	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatalf("error initiating Redis : %v", err)
		return
	}

	log.Println("Redis initiated")
}

func GetClient() *redis.Client {
	return client
}
