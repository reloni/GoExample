package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func setRedisString(c redis.Conn, key string, value string) error {
	_, err := c.Do("SET", key, value)
	if err != nil {
		log.Printf("Redis. Error setting key %s: %s", key, err.Error())
		return err
	}
	log.Printf("Redis. Set %s", key)

	return nil
}

func getRedisString(c redis.Conn, key string) (*string, error) {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		log.Printf("Redis. Error geting key %s: %s", key, err.Error())
		return nil, err
	}

	return &s, nil
}

func setRedisExpire(c redis.Conn, key string, value int) {
	_, err := c.Do("EXPIRE", key, value)
	if err != nil {
		log.Printf("Redis. Error seting expire key %s: %s", key, err.Error())
	}
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
}
