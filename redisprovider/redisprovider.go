package redisprovider

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var client *redis.Client

func LoadClient(host string, port string, pass string) error {
	log.Println("initiating Redis")

	address := fmt.Sprintf("%s:%s", host, port)

	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: pass,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatalf("error initiating Redis : %v", err)
		return err
	}

	log.Println("Redis initiated")

	return nil
}

func GetClient() *redis.Client {
	return client
}
