package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheImages interface {
	SetImageBytes(key string, value string)
	GetImageBytes(key string) string
	CheckExist(key string) bool
}

type Cache struct {
	CacheImages
}

func NewCache(cache *redis.Client, duration time.Duration) *Cache {
	return &Cache{
		CacheImages: NewCacheImages(cache, duration),
	}

}
