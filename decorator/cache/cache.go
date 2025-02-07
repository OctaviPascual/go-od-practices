package cache

type Cacher interface {
	Add(key string, value int) error
	Get(key string) (int, error)
	Delete(key string) error
}

type LRUCache struct{}

func (c *LRUCache) Add(_ string, _ int) error { return nil }

func (c *LRUCache) Get(_ string) (int, error) { return 0, nil }

func (c *LRUCache) Delete(_ string) error { return nil }

func NewLRUCache() *LRUCache { return &LRUCache{} }

type LFUCache struct{}

func (c *LFUCache) Add(_ string, _ int) error { return nil }

func (c *LFUCache) Get(_ string) (int, error) { return 0, nil }

func (c *LFUCache) Delete(_ string) error { return nil }

func NewLFUCache() *LFUCache { return &LFUCache{} }
