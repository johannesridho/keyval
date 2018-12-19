package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var redisClient *redis.Client

type Payload struct {
	Key   string
	Value string
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	router := mux.NewRouter()
	router.HandleFunc("/keyvals", setKeyval).Methods("POST")
	router.HandleFunc("/keyvals/{key}", getKeyVal).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
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

func getKeyVal(responseWriter http.ResponseWriter, request *http.Request) {
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