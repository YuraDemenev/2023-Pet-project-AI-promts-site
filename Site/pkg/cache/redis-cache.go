package cache

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) CacheImages {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (myCache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     myCache.host,
		Password: "",
		DB:       myCache.db,
	})
}

func (myCache *redisCache) SetImageBytes(key string, value string) {
	client := myCache.getClient()
	ctx := context.Background()

	var mutex = &sync.Mutex{}
	mutex.Lock()
	client.Set(ctx, key, value, time.Hour)
	mutex.Unlock()
}

func (myCache *redisCache) GetImageBytes(key string) string {
	client := myCache.getClient()
	ctx := context.Background()

	var mutex = &sync.RWMutex{}
	mutex.RLock()
	val, err := client.Get(ctx, key).Result()
	mutex.RUnlock()

	if err != nil {
		logrus.Fatalf("failed to get from redis value with key %s : %s", key, err.Error())
	}
	return val
}

func (myCache *redisCache) CheckExist(key string) bool {
	client := myCache.getClient()
	ctx := context.Background()

	check, err := client.Exists(ctx, key).Result()
	if err != nil {
		logrus.Fatalf("failed to get from redis value with key %s : %s", key, err.Error())
	}

	if check == 0 {
		return false
	} else {
		return true
	}
}
