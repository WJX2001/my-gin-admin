package gva_cache

import "time"

type Cache interface {
	Get(key string) (any, bool)
	Set(key string, value any, ttl time.Duration)
	SetDefault(key string, value any)
	Increment(key string, n int64) (int64, error)
	IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error)
	Exists(key string) bool
	Delete(key string)
}
