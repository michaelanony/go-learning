package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis()(err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:               "192.168.11.31:30002",
		Password:           "cbcbcb",

	})
	_,err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err !=nil{
		panic(err)
		return
	}
	fmt.Println("success")
}