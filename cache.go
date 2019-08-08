package poeapi

import (
	"container/list"
	"sync"
)

// cache is an in-memory LRU cache with a fixed size in terms of item count.
// It is based on https://github.com/hashicorp/golang-lru, combining the
// features of the simplelru package with the features of its threadsafe
// top-level package. Keys and values must be strings.
type cache struct {
	size      int
	evictList *list.List
	items     map[string]*list.Element
	lock      sync.RWMutex
}

type entry struct {
	key   string
	value string
}

func newCache(size int) (*cache, error) {
	if size <= 0 {
		return nil, ErrInvalidCacheSize
	}
	c := &cache{
		size:      size,
		evictList: list.New(),
		items:     make(map[string]*list.Element),
	}
	return c, nil
}

func (c *cache) Add(key, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
	}
	ent := &entry{key, value}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry
	if c.evictList.Len() > c.size {
		c.removeOldest()
	}
}

// Get retrieves an item from the cache. If the requested key is not in the
// cache, an empty string is returned.
func (c *cache) Get(key string) string {
	c.lock.Lock()
	defer c.lock.Unlock()

	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return ""
		}
		return ent.Value.(*entry).value
	}
	return ""
}

func (c *cache) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

func (c *cache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
}
