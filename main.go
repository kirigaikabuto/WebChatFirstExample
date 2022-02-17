package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	rdb *redis.Client
)

func main() {
	redisURL := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Print("Server starting at localhost:4444")
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
