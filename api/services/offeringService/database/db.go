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

var pool = newPool()
var keys []interface{}

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
	keys = append(keys, key)
	if err != nil {
		return err
	}

	return nil
}

func GetAll() (res []string) {
	client := pool.Get()
	defer client.Close()

	for _, key := range keys {
		value, _ := Get(key)
		res = append(res, value.(string)+" : "+key.(string))
	}

	return
}
