package poeapi

import (
	"container/list"
	"container/ring"
	"net"
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

// entry is an item in the cache's eviction list.
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

// Add adds an item to the cache.
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

// dnscache is an in-memory ring cache which caches IP addresses from DNS
// resolution for api.pathofexile.com. DNS can be a significant factor in
// request latency, and Go does not cache DNS resolution by default.
// TODO: Make the cache generic so it can be used for more than one host.
type dnscache struct {
	ips *ring.Ring
}

func (d *dnscache) getIP() (string, error) {
	ipval := d.ips.Value
	ip, ok := ipval.(string)
	if !ok {
		return "", ErrInvalidAddress
	}

	d.ips = d.ips.Next()
	return ip, nil
}

func newDNSCache(host string) (*dnscache, error) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}

	r := ring.New(len(addrs))
	for i := 0; i < r.Len(); i++ {
		r.Value = addrs[i]
		r = r.Next()
	}

	return &dnscache{ips: r}, nil
}
