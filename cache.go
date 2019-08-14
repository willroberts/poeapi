package poeapi

import (
	"container/list"
	"container/ring"
	"net"
	"sync"
)

// responsecache stores JSON responses from the API, storing them by URL. It is
// thread-safe and uses strings for both keys and values. It tracks recently
// used URLs deletes the oldest entries when maxSize is reached.
type responsecache struct {
	responses  map[string]*list.Element
	recenturls *list.List
	maxSize    int
	lock       sync.Mutex
}

type response struct {
	url  string
	body string
}

// Get retrieves a response from the cache.
func (c *responsecache) Get(url string) (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	resp, ok := c.responses[url]
	if !ok {
		return "", ErrNotFoundInCache
	}

	c.recenturls.MoveToFront(resp)
	return resp.Value.(*response).body, nil
}

// Set writes a response to the cache.
func (c *responsecache) Set(url, body string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if resp, ok := c.responses[url]; ok {
		c.recenturls.MoveToFront(resp)
		resp.Value.(*response).body = body
		return
	}
	c.responses[url] = c.recenturls.PushFront(&response{
		url:  url,
		body: body,
	})

	if c.recenturls.Len() > c.maxSize {
		oldest := c.recenturls.Back()
		c.recenturls.Remove(oldest)
		delete(c.responses, oldest.Value.(*response).url)
	}
}

func newResponseCache(maxSize int) (*responsecache, error) {
	if maxSize < 1 {
		return nil, ErrInvalidCacheSize
	}
	return &responsecache{
		responses:  make(map[string]*list.Element),
		recenturls: list.New(),
		maxSize:    maxSize,
	}, nil
}

// dnscache is an in-memory ring cache which caches IP addresses from DNS
// resolution for api.pathofexile.com. DNS can be a significant factor in
// request latency, and Go does not cache DNS resolution by default.
type dnscache struct {
	IPs map[string]*ring.Ring
}

// Get retrieves the least-recently-used IP address from the DNS cache. If there
// is no cache entry, a DNS resolution is performed before returning an IP.
func (d *dnscache) Get(host string) (string, error) {
	if _, ok := d.IPs[host]; !ok {
		if err := d.resolve(host); err != nil {
			return "", err
		}
	}

	ip, ok := d.IPs[host].Value.(string)
	if !ok {
		return "", ErrInvalidAddress
	}
	d.IPs[host] = d.IPs[host].Next()
	return ip, nil
}

func (d *dnscache) resolve(host string) error {
	addrs, err := net.LookupHost(host)
	if err != nil {
		return err
	}
	r := ring.New(len(addrs))
	for i := 0; i < r.Len(); i++ {
		r.Value = addrs[i]
		r = r.Next()
	}
	d.IPs[host] = r
	return nil
}

func newDNSCache() *dnscache {
	ips := make(map[string]*ring.Ring, 0)
	return &dnscache{IPs: ips}
}
