package cache

type CacheImages interface {
	SetImageBytes(key string, value string)
	GetImageBytes(key string) string
	CheckExist(key string) bool
}
