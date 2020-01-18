package main

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool

func initPool(address string,)  {
	pool = &redis.Pool{
		MaxIdle:8,
		MaxActive:0,
		IdleTimeout:100,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp",address,)
		},

	}
}