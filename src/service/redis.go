package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func setString(c redis.Conn, key string, value string) error {
	_, err := c.Do("SET", key, value)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return err
	}

	log.Printf("Set %s in Redis", key)

	return nil
}

// get executes the redis GET command
func getString(c redis.Conn, key string) (*string, error) {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}

	return &s, nil
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
}
