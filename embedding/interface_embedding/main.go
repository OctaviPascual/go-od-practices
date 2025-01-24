package main

import (
	"fmt"

	"github.com/OctaviPascual/go-testing-practices/embedding/cache"
)

type VerboseCache struct {
	cache.Cacher
	adds int
}

func (c *VerboseCache) Add(key string, value int) error {
	fmt.Printf("addding key %s to cache\n", key)
	c.adds++
	return c.Cacher.Add(key, value)
}

func (c *VerboseCache) numberOfAdds() int { return c.adds }

func emptyCache(c cache.Cacher) {
	_ = c.Delete("1")
	_ = c.Delete("2")
	_ = c.Delete("3")
}

func main() {
	verboseCache := &VerboseCache{Cacher: cache.NewLRUCache(), adds: 0}
	verboseCache = &VerboseCache{Cacher: cache.NewLFUCache(), adds: 0}

	_ = verboseCache.Add("1", 1)
	_ = verboseCache.Add("2", 2)

	_ = verboseCache.Delete("1")
	_, _ = verboseCache.Get("1")

	emptyCache(verboseCache)

	fmt.Printf("Total number of adds: %d", verboseCache.numberOfAdds())
}
