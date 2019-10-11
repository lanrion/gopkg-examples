package main

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
)

var cache *redis.Client
var once sync.Once

func GetCache() *redis.Client {
	once.Do(initCache)
	return cache
}

func initCache() {
	cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := cache.Ping().Err(); err != nil {
		log.Println("initCache")
	}
}
