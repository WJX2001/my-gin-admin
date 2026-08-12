package gva_cache

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"time"
)

type memoryCache struct {
	c             local_cache.Cache
	defaultExpire time.Duration
}

// NewMemoryCache 创建内存后端，defaultExpire 为 SetDefault 的默认过期时间
func NewMemoryCache(defaultExpire time.Duration) Cache {
	return &memoryCache{
		c: local_cache.NewCache(
			local_cache.SetDefaultExpire(defaultExpire),
			local_cache.SetCapture(func(string, interface{}) {}), // 禁用默认 stdout 打印
		),
		defaultExpire: defaultExpire,
	}
}

func (m *memoryCache) Get(key string) (any, bool) {
	return m.c.Get(key)
}

func (m *memoryCache) Set(key string, value any, ttl time.Duration) {
	m.c.Set(key, value, ttl)
}

func (m *memoryCache) SetDefault(key string, value any) {
	m.c.SetDefault(key, value)
}

// increment 内部统一逻辑：key 不存在则按 int64(n) 建立计数（带 ttl），存在则 IncrementInt64。
func (m *memoryCache) increment(key string, n int64, ttl time.Duration) (int64, error) {
	if _, ok := m.c.Get(key); !ok {
		m.c.Set(key, n, ttl)
		return n, nil
	}
	return m.c.IncrementInt64(key, n)
}

// increment 内部统一逻辑：key 不存在则按 int64(n) 建立计数（带 ttl），存在则 IncrementInt64。
func (m *memoryCache) Increment(key string, n int64) (int64, error) {
	return m.increment(key, n, m.defaultExpire)
}

func (m *memoryCache) IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error) {
	return m.increment(key, n, ttl)
}

func (m *memoryCache) Exists(key string) bool {
	_, ok := m.c.Get(key)
	return ok
}

func (m *memoryCache) Delete(key string) {
	m.c.Delete(key)
}
