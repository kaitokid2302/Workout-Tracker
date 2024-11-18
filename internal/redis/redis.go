package redis

import "github.com/redis/go-redis/v9"

func RedisInit() *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // no password set
		DB:       0,                // default DB
	})
	return Client
}
