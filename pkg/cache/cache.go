package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cacher interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{})
}

type Cache struct {
	cache             *cache.Cache
	defaultExpiration time.Duration
}

func New(defaultExpiration, cleanupInterval int) *Cache {
	expiration := time.Duration(defaultExpiration) * time.Minute

	return &Cache{
		cache: cache.New(
			expiration,
			time.Duration(cleanupInterval)*time.Minute),
		defaultExpiration: expiration,
	}
}

func (c *Cache) Get(k string) (interface{}, bool) {
	return c.cache.Get(k)
}

func (c *Cache) Set(k string, x interface{}) {
	c.cache.Set(k, x, c.defaultExpiration)
}
