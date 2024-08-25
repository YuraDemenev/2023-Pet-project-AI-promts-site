package cache

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisCache struct {
	cache    *redis.Client
	duration time.Duration
}

func NewCacheImages(cache *redis.Client, duration time.Duration) *RedisCache {
	return &RedisCache{cache: cache, duration: duration}
}

func (myCache *RedisCache) SetImageBytes(key string, value string) {
	ctx := context.Background()

	var mutex = &sync.Mutex{}
	mutex.Lock()
	myCache.cache.Set(ctx, key, value, time.Hour)
	mutex.Unlock()
}

func (myCache *RedisCache) GetImageBytes(key string) string {
	ctx := context.Background()

	var mutex = &sync.RWMutex{}
	mutex.RLock()
	val, err := myCache.cache.Get(ctx, key).Result()
	mutex.RUnlock()

	if err != nil {
		logrus.Fatalf("failed to get from redis value with key %s : %s", key, err.Error())
	}
	return val
}

func (myCache *RedisCache) CheckExist(key string) bool {
	ctx := context.Background()

	check, err := myCache.cache.Exists(ctx, key).Result()
	if err != nil {
		logrus.Fatalf("failed to get from redis value with key %s : %s", key, err.Error())
	}

	if check == 0 {
		return false
	} else {
		return true
	}
}
