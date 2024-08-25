package cache

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	Host string
	DB   int
}

func NewRedisCache(redisConfig RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: "",
		DB:       redisConfig.DB,
	})
	return client
}
