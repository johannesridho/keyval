package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/johannesridho/keyval/config"
	"github.com/johannesridho/keyval/redisprovider"
	"log"
	"net/http"
)

var redisClient *redis.Client

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	config.LoadEnv()

	err := redisprovider.LoadClient(config.RedisHost, config.RedisPort, config.RedisPassword)
	if err != nil {
		log.Fatalf("error initiating Redis : %v", err)
	}

	redisClient = redisprovider.GetClient()

	address := fmt.Sprintf("%s:%s", config.Host, config.Port)

	router := mux.NewRouter()
	router.HandleFunc("/keyvals", setKeyval).Methods("POST")
	router.HandleFunc("/keyvals/{key}", getKeyval).Methods("GET")
	log.Fatal(http.ListenAndServe(address, router))
}

func setKeyval(responseWriter http.ResponseWriter, request *http.Request) {
	var payload Payload
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&payload); err != nil {
		createErrorResponse(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer request.Body.Close()

	err := redisClient.Set(payload.Key, payload.Value, 0).Err()
	if err != nil {
		panic(err)
	}

	createJsonResponse(responseWriter, http.StatusOK, payload)
}

func getKeyval(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	key := vars["key"]
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		createErrorResponse(responseWriter, http.StatusBadRequest, "Key does not exist")
		return
	} else if err != nil {
		panic(err)
	}

	createJsonResponse(responseWriter, http.StatusOK, Payload{Key: key, Value: value})
}

func createErrorResponse(responseWriter http.ResponseWriter, code int, message string) {
	createJsonResponse(responseWriter, code, map[string]string{"error": message})
}

func createJsonResponse(responseWriter http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}
