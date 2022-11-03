package database

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// creating a new poll for redis connections
var pool = newPool()

func Get(key interface{}) (interface{}, error) {
	client := pool.Get()
	defer client.Close()

	value, err := client.Do("GET", key)
	if err != nil {
		return nil, errors.New("cant find the key")
	}

	return value, nil
}

func Set(key, value interface{}) error {
	client := pool.Get()
	defer client.Close()

	value, err := client.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// GetAll gets all keys by regex command
func GetAll() (interface{}, error) {
	client := pool.Get()
	defer client.Close()

	value, err := client.Do("KEYS *")
	if err != nil {
		return nil, err
	}

	return value, nil
}
